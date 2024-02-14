package main

import (
	holfilHdls "b1staticcontentfprint/handlers"
	flHdls "c1jsonmessages/handlers"
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

	// Use handlers from FileHandlers.go
	http.HandleFunc("/api/customers", flHdls.GetCustomersInJsonHandler)
	http.HandleFunc("/api/customers/add", flHdls.AddCustomersFromJsonHandler)

	fmt.Printf("Starting Web Server at http://localhost%s\n", addr)

	go func() {
		log.Fatal(s.ListenAndServe())
	}()

	fmt.Println("Server started, press <Enter> to shutdown")
	fmt.Scanln()
	s.Shutdown(context.Background())
	fmt.Println("Server stopped")
}
