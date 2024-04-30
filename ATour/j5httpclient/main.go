package main

import (
	utl "autilities"
	"bufio"
	"net/http"
)

var header = utl.Header{}

/*
The Go standard library comes with excellent support for HTTP clients and servers in the net/http package.
*/
func main() {
	/*
	 */
	header.DisplayHeader("Showing HTTP Client")

	// Issue an HTTP GET request to a server. http.Get is a convenient shortcut around creating an http.Client object and calling its Get method; it uses the http.DefaultClient object which has useful default settings.

	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Print the HTTP response status.

	utl.PLine("Response status:", resp.Status)

	// Print the first 5 lines of the response body.

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		utl.PLine(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
