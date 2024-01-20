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

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Welcome to Customer Service Web Service.\n")
	}

	// http://localhost:8080
	http.HandleFunc("/", helloHandler)

	// http://localhost:8080/api
	http.HandleFunc("/api", helloHandler)

	var getUrlHandlerFunc http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.URL.String())
	}

	// http://localhost:8080/url/ or http://localhost:8080/url/something or http://localhost:8080/url/something/else
	http.HandleFunc("/url/", getUrlHandlerFunc)

	fmt.Printf("Starting Web Server at http://localhost%s\n", addr)

	go func() {
		log.Fatal(s.ListenAndServe())
	}()

	fmt.Println("Server started, press <Enter> to shutdown")
	fmt.Scanln()
	s.Shutdown(context.Background())
	fmt.Println("Server stopped")
}
