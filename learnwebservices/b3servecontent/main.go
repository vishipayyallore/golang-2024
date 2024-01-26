// main.go

package main

import (
	"b3servecontent/handlers"
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

	// Use handlers from filehandlers.go
	http.HandleFunc("/", handlers.HelloHandler)
	http.HandleFunc("/api", handlers.HelloHandler)
	http.HandleFunc("/url/", handlers.GetUrlHandlerFunc)

	http.HandleFunc("/api/getcustomerdata", handlers.GetCustomerDataHandler)
	http.HandleFunc("/api/getcustomerdatav1", handlers.ServeFileHandler)
	http.HandleFunc("/api/getcustomerdatav2", handlers.ServeContentHandler)

	fmt.Printf("Starting Web Server at http://localhost%s\n", addr)

	go func() {
		log.Fatal(s.ListenAndServe())
	}()

	fmt.Println("Server started, press <Enter> to shutdown")
	fmt.Scanln()
	s.Shutdown(context.Background())
	fmt.Println("Server stopped")
}
