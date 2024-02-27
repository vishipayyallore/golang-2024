// main.go

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

	// GET http://localhost:8081
	http.HandleFunc("/", pHdls.ServerHomeHtml)

	// GET http://localhost:8081/api
	http.HandleFunc("/api", pHdls.ServerHomeHtml)

	// GET http://localhost:8081/api/products
	http.HandleFunc("/api/products", pHdls.GetAllProductsHandler)

	// GET http://localhost:8081/api/products/?id=1
	http.HandleFunc("/api/products/", pHdls.GetAllProductByQueryStringHandler)

	// GET http://localhost:8081/api/products-qs?id=2
	http.HandleFunc("/api/products-qs", pHdls.GetAllProductByQueryStringHandler)

	// GET http://localhost:8081/api/products-ssplit/3
	http.HandleFunc("/api/products-ssplit/", pHdls.GetAllProductByRouteParameterHandler)

	fmt.Printf("Starting Web Server at http://localhost%s\n", addr)

	go func() {
		log.Fatal(s.ListenAndServe())
	}()

	fmt.Println("Server started, press <Enter> to shutdown")
	fmt.Scanln()
	s.Shutdown(context.Background())
	fmt.Println("Server stopped")
}
