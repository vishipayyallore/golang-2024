// main.go

package main

import (
	holfilHdls "b1staticcontentfprint/handlers"
	fileHdls "b2staticcontentservefile/handlers"
	fileHdls2 "b3servecontent/handlers"
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
	http.HandleFunc("/", holfilHdls.HelloHandler)
	http.HandleFunc("/api", holfilHdls.HelloHandler)
	http.HandleFunc("/url/", holfilHdls.GetUrlHandlerFunc)

	// Use handlers from FileHandlers.go
	http.HandleFunc("/api/getcustomerdatav1", holfilHdls.GetCustomerDataHandlerv1)
	http.HandleFunc("/api/getcustomerdatav2", holfilHdls.GetCustomerDataHandlerv2)
	http.HandleFunc("/api/getcustomerdatav3", fileHdls.ServeFileHandler)
	http.HandleFunc("/api/getcustomerdatav4", fileHdls2.ServeContentHandler)

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
