// main.go

package main

import (
	"context"
	pHdls "d1basicrouting/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	// GET http://localhost:8082
	r.HandleFunc("/", pHdls.ServerHomeHtml)

	// GET http://localhost:8082/api
	r.HandleFunc("/api", pHdls.ServerHomeHtml)

	// GET http://localhost:8082/api/products
	r.HandleFunc("/api/products", pHdls.GetAllProductsHandler)

	// GET http://localhost:8082/api/products/?id=1
	r.HandleFunc("/api/products/", pHdls.GetAllProductByQueryStringHandler)

	// GET http://localhost:8082/api/products-qs?id=2
	r.HandleFunc("/api/products-qs", pHdls.GetAllProductByQueryStringHandler)

	// GET http://localhost:8082/api/products-ssplit/3
	r.HandleFunc("/api/products-ssplit/", pHdls.GetAllProductByRouteParameterHandlerStrSplit)

	// GET http://localhost:8082/api/products-regexp/4
	r.HandleFunc("/api/products-regexp/", pHdls.GetAllProductByRouteParameterHandlerRegExp)

	http.Handle("/", r)

	addr := ":8082"
	s := http.Server{
		Addr: addr,
	}

	fmt.Printf("Starting Web Server at http://localhost%s\n", addr)

	go func() {
		log.Fatal(s.ListenAndServe())
	}()

	fmt.Println("Server started, press <Enter> to shutdown")
	fmt.Scanln()
	s.Shutdown(context.Background())
	fmt.Println("Server stopped")
}
