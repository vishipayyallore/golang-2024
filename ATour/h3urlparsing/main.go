package main

import (
	utl "autilities"
	"net/url"
)

var header = utl.Header{}

/*
URLs provide a uniform way to locate resources. Here’s how to parse URLs in Go.
*/
func main() {
	/*
	 */
	header.DisplayHeader("Showing URL Parsing")

	// We’ll parse this example URL, which includes a scheme, authentication info, host, port, path, query params, and query fragment.
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// Parse the URL and ensure there are no errors.
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// Accessing the scheme is straightforward.
	utl.PLine("Scheme: ", u.Scheme)
}
