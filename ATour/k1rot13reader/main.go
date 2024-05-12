package main

import (
	utl "autilities"
	"fmt"
	"io"
	"os"
	"strings"
)

var header = utl.Header{}

/*
 */
func main() {
	/*
	 */
	header.DisplayHeader("Showing Exercise: rot13Reader")

	s := strings.NewReader("This is a Sample line!")
	fmt.Println("Original: ", s)
	r := &rot13Reader{s}
	io.Copy(os.Stdout, r)
}

type rot13Reader struct {
	r io.Reader
}

func (rr *rot13Reader) Read(p []byte) (int, error) {
	n, err := rr.r.Read(p)
	if err != nil {
		return n, err
	}

	for i := 0; i < n; i++ {
		c := p[i]
		if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
			// Apply ROT13 to alphabetical characters
			switch {
			case (c >= 'A' && c <= 'Z'):
				p[i] = 'A' + (c-'A'+13)%26
			case (c >= 'a' && c <= 'z'):
				p[i] = 'a' + (c-'a'+13)%26
			}
		}
	}

	return n, nil
}
