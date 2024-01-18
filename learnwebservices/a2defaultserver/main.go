package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	addr := ":8080"

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Welcome to Customer Service Web Service.\n")
	}

	http.HandleFunc("/", helloHandler)

	http.HandleFunc("/api", helloHandler)

	fmt.Printf("Starting Web Server at http://localhost%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
