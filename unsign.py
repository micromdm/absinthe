import objc
from Foundation import NSBundle

Security_bundle = NSBundle.bundleWithIdentifier_('com.apple.security')

CMSDecoderRef = objc.createOpaquePointerType("CMSDecoderRef", b"^{CMSDecoder}", None)

functions = [('CMSDecoderCreate',          b'io^^{CMSDecoder}'),
             ('CMSDecoderUpdateMessage',   b'i^{CMSDecoder}*I'),
             ('CMSDecoderFinalizeMessage', b'i^{CMSDecoder}',),
             ('CMSDecoderCopyContent',     b'i^{CMSDecoder}o^^{__CFData}')]

objc.loadBundleFunctions(Security_bundle, globals(), functions)

f = open('cms.txt', 'rb')
signed_plist = f.read()
f.close()

err, decoder = CMSDecoderCreate(None)
print err
err = CMSDecoderUpdateMessage(decoder, signed_plist, len(signed_plist))
print err
err = CMSDecoderFinalizeMessage(decoder)
print err
err, unsigned_data = CMSDecoderCopyContent(decoder, None)

# plist_bytes = unsigned_data.bytes().tobytes()
# print plist_bytes
