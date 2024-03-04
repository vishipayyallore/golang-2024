// main.go

package main

import (
	"context"
	pHdls "d1basicrouting/handlers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {

	r := mux.NewRouter()

	// Use logrus logger
	var log = logrus.New()

	// Set log level
	log.SetLevel(logrus.DebugLevel)

	// GET http://localhost:8081
	r.HandleFunc("/", pHdls.ServerHomeHtml)

	// GET http://localhost:8081/api
	r.HandleFunc("/api", pHdls.ServerHomeHtml)

	// GET http://localhost:8081/api/products
	r.HandleFunc("/api/products", pHdls.GetAllProductsHandler)

	// GET http://localhost:8081/api/products/?id=1
	r.HandleFunc("/api/products/", pHdls.GetAllProductByQueryStringHandler)

	// GET http://localhost:8081/api/products-qs?id=2
	r.HandleFunc("/api/products-qs", pHdls.GetAllProductByQueryStringHandler)

	// GET http://localhost:8081/api/products-ssplit/3
	r.HandleFunc("/api/products-ssplit/", pHdls.GetAllProductByRouteParameterHandler)

	// GET http://localhost:8081/api/products-regexp/4
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
