// main.go

package main

import (
	holfilHdls "b1staticcontentfprint/handlers"
	"b2staticcontentservefile/handlers"
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
	http.HandleFunc("/", holfilHdls.HelloHandler)
	http.HandleFunc("/api", holfilHdls.HelloHandler)
	http.HandleFunc("/url/", holfilHdls.GetUrlHandlerFunc)

	http.HandleFunc("/api/getcustomerdatav1", holfilHdls.GetCustomerDataHandlerv1)
	http.HandleFunc("/api/getcustomerdatav2", holfilHdls.GetCustomerDataHandlerv2)
	http.HandleFunc("/api/getcustomerdatav3", handlers.ServeFileHandler)

	fmt.Printf("Starting Web Server at http://localhost%s\n", addr)

	go func() {
		log.Fatal(s.ListenAndServe())
	}()

	fmt.Println("Server started, press <Enter> to shutdown")
	fmt.Scanln()
	s.Shutdown(context.Background())
	fmt.Println("Server stopped")
}
