// main.go

package main

import (
	hloHdlers "b2staticcontentservefile/handlers"
	"b3servecontent/handlers"
	fSrvHdlers "b4fileserver/handlers"
	"context"
	"fmt"
	"log"
	"net/http"
)

const (
	fileServerRoute   string = "/downloads/"
	downloadsFilePath string = "../downloads"
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

	// Use handlers from FileHandlers.go
	http.HandleFunc("/api/getcustomerdata", handlers.GetCustomerDataHandler)
	http.HandleFunc("/api/getcustomerdatav1", handlers.ServeFileHandler)
	http.HandleFunc("/api/getcustomerdatav2", handlers.ServeContentHandler)

	// File Server
	http.Handle(fileServerRoute, fSrvHdlers.GetFileServerHandlerFunc(fileServerRoute, downloadsFilePath))

	fmt.Printf("Starting Web Server at http://localhost%s\n", addr)

	go func() {
		log.Fatal(s.ListenAndServe())
	}()

	fmt.Println("Server started, press <Enter> to shutdown")
	fmt.Scanln()
	s.Shutdown(context.Background())
	fmt.Println("Server stopped")
}

// http.Handle("/downloads/", http.StripPrefix("/downloads/", http.FileServer(http.Dir("../downloads"))))
