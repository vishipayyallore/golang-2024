package main

import (
	utl "autilities"
	"crypto/sha256"
)

var header = utl.Header{}

/*
SHA256 hashes are frequently used to compute short identities for binary or text blobs. For example, TLS/SSL certificates use SHA256 to
compute a certificate’s signature. Here’s how to compute SHA256 hashes in Go.
*/
func main() {
	/*
		Go implements several hash functions in various crypto/* packages.
	*/
	header.DisplayHeader("Showing SHA256 Hashes")

	s := "sha256 this string"

	// Here we start with a new hash.
	h := sha256.New()
	utl.PLine("String: ", s)
	utl.PLine("Hash: ", h)

	// Write expects bytes. If you have a string s, use []byte(s) to coerce it to bytes.
	h.Write([]byte(s))

	// This gets the finalized hash result as a byte slice. The argument to Sum can be used to append to an existing byte slice: it usually isn’t needed.
	bs := h.Sum(nil)
	utl.PLine("SHA256 Hash: ", bs)

	// SHA1 values are often printed in hex, for example in git commits. Use the %x format verb to convert a hash results to a hex string.
	utl.PFmted("Hex String: %x\n", bs)

	// This gets the finalized hash result as a byte slice. The argument to Sum can be used to append to an existing byte slice: it usually isn’t needed.
}
