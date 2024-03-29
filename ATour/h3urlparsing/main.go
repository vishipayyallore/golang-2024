package main

import (
	utl "autilities"
	"net"
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

	// User contains all authentication info; call Username and Password on this for individual values.
	utl.PLine("User: ", u.User)
	utl.PLine("Username: ", u.User.Username())
	p, _ := u.User.Password()
	utl.PLine("Password: ", p)

	// The Host contains both the hostname and the port, if present. Use SplitHostPort to extract them.
	utl.PLine("Host: ", u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	utl.PLine("Host: ", host)
	utl.PLine("Port: ", port)

	// Here we extract the path and the fragment after the #.
	utl.PLine("Path: ", u.Path)
	utl.PLine("Fragment: ", u.Fragment)

	// To get query params in a string of k=v format, use RawQuery. You can also parse query params into a map. The parsed query param maps are from strings to slices of strings, so index into [0] if you only want the first value.
	utl.PLine("RawQuery: ", u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	utl.PLine("Query Params: ", m)
	utl.PLine("Query Param k: ", m["k"][0])
}
