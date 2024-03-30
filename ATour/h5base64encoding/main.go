package main

import (
	utl "autilities"
	b64 "encoding/base64"
)

var header = utl.Header{}

/*
Go provides built-in support for base64 encoding/decoding. The encoding/base64 package implements base64 encoding as specified by RFC 4648.
*/
func main() {
	/*
		This syntax imports the encoding/base64 package with the b64 name instead of the default base64.
	*/
	header.DisplayHeader("Showing Base64 Encoding")

	// Here’s the string we’ll encode/decode.
	data := "abc123!?$*&()'-=@~"
	utl.PLine("Original Data: ", data)

	// Go supports both standard and URL-compatible base64. Here’s how to encode using the standard encoder. The encoder requires a []byte
	// so we convert our string to that type.
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	utl.PLine("Standard Encoded: ", sEnc)

	// Decoding may return an error, which you can check if you don’t already know the input to be well-formed.
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	utl.PLine("Standard Decoded: ", string(sDec))

	data = "https://gobyexample.com/base64-encoding"
	utl.PLine("\nOriginal Data: ", data)
	// This encodes/decodes using a URL-compatible base64 format.
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	utl.PLine("URL Encoded: ", uEnc)

	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	utl.PLine("URL Decoded: ", string(uDec))
}
