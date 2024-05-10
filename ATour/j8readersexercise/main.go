package main

import (
	utl "autilities"

	"golang.org/x/tour/reader"
)

var header = utl.Header{}

/*
 */
func main() {
	/*
	 */
	header.DisplayHeader("Showing Exercise: Readers")

	reader.Validate(MyReader{})
}

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	// Fill the byte slice with 'A' characters.
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}
