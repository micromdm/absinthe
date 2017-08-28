Disclaimer: This repo is just some notes/experimental code.

These repo sets up a local DEP server proxy.

First, we set up a local server which can respond to the same URL endpoints as the DEP server.
Next, run `./configure-dep.sh` script. The values defined in `configure-dep.sh` configure the mac to make requests to the local Go proxy.
Finally, run `dep nag` which will make the DEP profile request (through the proxy). 

You can manipulate the Go code to intercept the profile returned by the server. Look at `profileHandler` in `main.go`
