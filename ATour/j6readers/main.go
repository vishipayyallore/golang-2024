package main

import (
	utl "autilities"
	"fmt"
	"io"
	"strings"
)

var header = utl.Header{}

/*
The io package specifies the io.Reader interface, which represents the read end of a stream of data.
*/
func main() {
	/*
		The Go standard library contains many implementations of this interface, including files, network connections, compressors,
		ciphers, and others.
	*/
	header.DisplayHeader("Showing Readers")

	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		// Read populates the given byte slice with data and returns the number of bytes populated and an error value. It returns an io.EOF error when the stream ends.
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
