// main.go

package main

import (
	fileHdlers "b1staticcontentfprint/handlers"
	hloHdlers "b1staticcontentfprint/handlers"
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
	http.HandleFunc("/", hloHdlers.HelloHandler)
	http.HandleFunc("/api", hloHdlers.HelloHandler)
	http.HandleFunc("/url/", hloHdlers.GetUrlHandlerFunc)

	http.HandleFunc("/api/getcustomerdata", fileHdlers.GetCustomerDataHandlerv1)
	http.HandleFunc("/api/getcustomerdatav1", handlers.ServeFileHandler)

	fmt.Printf("Starting Web Server at http://localhost%s\n", addr)

	go func() {
		log.Fatal(s.ListenAndServe())
	}()

	fmt.Println("Server started, press <Enter> to shutdown")
	fmt.Scanln()
	s.Shutdown(context.Background())
	fmt.Println("Server stopped")
}
