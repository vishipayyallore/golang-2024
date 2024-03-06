// main.go

package main

import (
	"context"
	pHdls "d1basicrouting/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {

	r := chi.NewRouter()

	// GET http://localhost:8083
	r.Get("/", pHdls.ServerHomeHtml)

	// GET http://localhost:8083/api
	r.Get("/api", pHdls.ServerHomeHtml)

	// GET http://localhost:8083/api/products
	r.Get("/api/products", pHdls.GetAllProductsHandler)

	// GET http://localhost:8083/api/products/?id=1
	r.Get("/api/products/", pHdls.GetAllProductByQueryStringHandler)

	// GET http://localhost:8083/api/products-qs?id=2
	r.Get("/api/products-qs", pHdls.GetAllProductByQueryStringHandler)

	// GET http://localhost:8083/api/products-ssplit/3
	r.Get("/api/products/{id:[0-9]+}", pHdls.GetAllProductByRouteParameterHandlerGoChi)

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
