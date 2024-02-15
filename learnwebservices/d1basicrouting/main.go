package main

import (
	"context"
	pHdls "d1basicrouting/handlers"
	"fmt"
	"log"
	"net/http"
)

func main() {

	addr := ":8081"
	s := http.Server{
		Addr: addr,
	}

	http.HandleFunc("/api/products", pHdls.GetAllProductsHandler)
	http.HandleFunc("/api/products/", pHdls.GetAllProductByQueryHandler)

	fmt.Printf("Starting Web Server at http://localhost%s\n", addr)

	go func() {
		log.Fatal(s.ListenAndServe())
	}()

	fmt.Println("Server started, press <Enter> to shutdown")
	fmt.Scanln()
	s.Shutdown(context.Background())
	fmt.Println("Server stopped")
}
