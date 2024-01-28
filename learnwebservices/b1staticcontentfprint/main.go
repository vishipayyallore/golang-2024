package main

import (
	"b1staticcontentfprint/handlers"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	addr := ":8080"
	s := http.Server{
		Addr: addr,
	}

	// http://localhost:8080
	http.HandleFunc("/", handlers.HelloHandler)

	// http://localhost:8080/api
	http.HandleFunc("/api", handlers.HelloHandler)

	// http://localhost:8080/url/ or http://localhost:8080/url/something or http://localhost:8080/url/something/else
	http.HandleFunc("/url/", handlers.GetUrlHandlerFunc)

	filePath := "../data/customers.csv"
	getCustomerDataHandler := func(w http.ResponseWriter, req *http.Request) {
		customerFile, err := os.Open(filePath)
		if err != nil {
			log.Fatal(err)
		}
		defer customerFile.Close()

		// Method 1
		data, err := io.ReadAll(customerFile)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprint(w, string(data))

		// Method 2
		// io.Copy(w, customerFile)
	}

	// http://localhost:8080/api/getcustomerdata
	http.HandleFunc("/api/getcustomerdata", getCustomerDataHandler)

	fmt.Printf("Starting Web Server at http://localhost%s\n", addr)

	go func() {
		log.Fatal(s.ListenAndServe())
	}()

	fmt.Println("Server started, press <Enter> to shutdown")
	fmt.Scanln()
	s.Shutdown(context.Background())
	fmt.Println("Server stopped")
}
