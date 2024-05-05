package main

import (
	utl "autilities"
	"fmt"
	"net/http"
)

var header = utl.Header{}

/*
Writing a basic HTTP server is easy using the net/http package.
*/
func main() {
	/*
	 */
	header.DisplayHeader("Showing HTTP Server")

	// We register our handlers on server routes using the http.HandleFunc convenience function. It sets up the default router in
	// the net/http package and takes a function as an argument.

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	// Finally, we call the ListenAndServe with the port and a handler. nil tells it to use the default router we’ve just set up.
	fmt.Println("Starting HTTP Server on port 8090. \nPlease click the below link to see the output. \n\thttp://localhost:8090/hello \n\thttp://localhost:8090/headers")
	http.ListenAndServe(":8090", nil)
}

/*
A fundamental concept in net/http servers is handlers. A handler is an object implementing the http.Handler interface. A common way to
write a handler is by using the http.HandlerFunc adapter on functions with the appropriate signature. Functions serving as handlers
take a http.ResponseWriter and a http.Request as arguments. The response writer is used to fill in the HTTP response. Here our
simple response is just “hello\n”.
*/
func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello from HTTP Server\n")
}

/*
This handler does something a little more sophisticated by reading all the HTTP request headers and echoing them into the response body.
*/
func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}
