package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	addr := ":8080"
	s := http.Server{
		Addr: addr,
	}

	var getUrlHandlerFunc http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.URL.String())
	}

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Welcome to Customer Service Web Service.\n")
	}

	http.HandleFunc("/", helloHandler) // http://localhost:8080

	http.HandleFunc("/api", helloHandler) // http://localhost:8080/api

	http.HandleFunc("/url/", getUrlHandlerFunc) // http://localhost:8080/url/

	fmt.Printf("Starting Web Server at http://localhost%s\n", addr)

	go func() {
		log.Fatal(s.ListenAndServe())
	}()

	fmt.Println("Server started, press <Enter> to shutdown")
	fmt.Scanln()
	s.Shutdown(context.Background())
	fmt.Println("Server stopped")
}
