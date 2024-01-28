package main

import (
	"b1staticcontentfprint/handlers"
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {

	addr := ":8080"
	s := http.Server{
		Addr: addr,
	}

	// Use handlers from HelloHandlers.go
	// http://localhost:8080
	http.HandleFunc("/", handlers.HelloHandler)
	// http://localhost:8080/api
	http.HandleFunc("/api", handlers.HelloHandler)
	// http://localhost:8080/url/ or http://localhost:8080/url/something or http://localhost:8080/url/something/else
	http.HandleFunc("/url/", handlers.GetUrlHandlerFunc)

	// Use handlers from FileHandlers.go
	// http://localhost:8080/api/getcustomerdata
	http.HandleFunc("/api/getcustomerdatav1", handlers.GetCustomerDataHandlerv1)
	http.HandleFunc("/api/getcustomerdatav2", handlers.GetCustomerDataHandlerv2)

	fmt.Printf("Starting Web Server at http://localhost%s\n", addr)

	go func() {
		log.Fatal(s.ListenAndServe())
	}()

	fmt.Println("Server started, press <Enter> to shutdown")
	fmt.Scanln()
	s.Shutdown(context.Background())
	fmt.Println("Server stopped")
}
