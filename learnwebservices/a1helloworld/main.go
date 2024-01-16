package main

import (
	"fmt"
	"io"
	"net/http"
)

// Simple Http HandleFunc that prints "Hello World" to the browser

func main() {

	addr := ":8080"

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello Go Lang Web Service World !!\n")
	}

	http.HandleFunc("/", helloHandler)

	fmt.Printf("Starting Web Server at %s\n", addr)
	http.ListenAndServe(addr, nil)
}
