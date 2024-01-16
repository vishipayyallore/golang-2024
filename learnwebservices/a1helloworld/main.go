package main

import (
	"fmt"
	"net/http"
)

// Simple Http HandleFunc that prints "Hello World" to the browser

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Go Lang Web Service World !!")
	})

	http.ListenAndServe(":8080", nil)
}
