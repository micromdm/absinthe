package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

func debugHTTPmiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body := io.TeeReader(r.Body, os.Stderr)
		r.Body = ioutil.NopCloser(body)
		out, err := httputil.DumpRequest(r, true)
		if err != nil {
			log.Println(err)
		}
		fmt.Println("")
		fmt.Println(string(out))
		fmt.Println("")
		next.ServeHTTP(w, r)
	})
}

// dumps request and request body to stdout
func sessionHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Post("https://iprofiles.apple.com/session", r.Header.Get("Content-Type"), r.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	io.Copy(w, resp.Body)
	return

}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	rbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.Post("https://iprofiles.apple.com/macProfile", r.Header.Get("Content-Type"), ioutil.NopCloser(bytes.NewReader(rbody)))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body := io.TeeReader(resp.Body, os.Stderr)
	io.Copy(w, body)
	fmt.Println("")
	return
}

// dumps request and request body to stdout
// serves 'certificate.cer' as a response
func certHandler(w http.ResponseWriter, r *http.Request) {
	cert, err := ioutil.ReadFile("certificate.cer")
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(200)
	w.Write(cert)
}

func main() {
	http.Handle("/session", http.HandlerFunc(sessionHandler))
	http.Handle("/profile", debugHTTPmiddleware(http.HandlerFunc(profileHandler)))
	http.HandleFunc("/", certHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
