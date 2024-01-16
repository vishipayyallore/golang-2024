package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	addr := ":8080"

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Welcome to Customer Service Web Service.\n")
	}

	http.HandleFunc("/", helloHandler)

	http.HandleFunc("/api", helloHandler)

	fmt.Printf("Starting Web Server at %s\n", addr)
	http.ListenAndServe(addr, nil)
}
