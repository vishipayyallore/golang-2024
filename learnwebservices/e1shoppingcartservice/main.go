package main

import (
	"context"
	ent "e1shoppingcartservice/entities"
	svc "e1shoppingcartservice/services"
	utls "e1shoppingcartservice/utilities"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {
	cs := createCustomerService()
	ps := createProductService()
	scs := svc.CreateShoppingCartService()

	go func() {
		cs.ListenAndServe()
	}()

	go func() {
		ps.ListenAndServe()
	}()

	go func() {
		scs.ListenAndServe()
	}()

	time.Sleep(1 * time.Second)
	res, err := http.Get("http://localhost:5005/api/carts")
	if err != nil {
		log.Fatal(err)
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Response: ", string(data))

	fmt.Println("Services started, press <Enter> to shutdown")
	fmt.Scanln()
	cs.Shutdown(context.Background())
	ps.Shutdown(context.Background())
	scs.Shutdown(context.Background())
	fmt.Println("Services stopped")
}

func createProductService() *http.Server {

	mux := http.NewServeMux()

	mux.HandleFunc("/api/products", func(w http.ResponseWriter, r *http.Request) {
		data, err := json.Marshal(ent.Products)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Add("Content-Type", "application/json")

		w.Write(data)
	})

	pattern := regexp.MustCompile(`^\/api/products\/(\d+?)$`)
	mux.HandleFunc("/api/products/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request: ", r.URL.Path)
		matches := pattern.FindStringSubmatch(r.URL.Path)
		fmt.Println("Matches: ", matches)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(matches[1])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		for _, p := range ent.Products {
			if id == p.ID {
				data, err := json.Marshal(p)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				w.Header().Add("Content-Type", "application/json")
				w.Write(data)
				return
			}
		}
		w.WriteHeader(http.StatusNotFound)

	})

	s := http.Server{
		Addr:    ":4000",
		Handler: mux,
	}

	return &s
}

func createCustomerService() *http.Server {

	f, err := os.Open("customers.csv")
	if err != nil {
		log.Fatal(err)
	}
	customers, err := utls.ReadCustomers(f)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/api/customers", func(w http.ResponseWriter, r *http.Request) {

		data, err := json.Marshal(customers)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Add("content-type", "application/json")
		w.Write(data)

	})

	pattern := regexp.MustCompile(`^\/api/customers\/(\d+?)$`)
	mux.HandleFunc("/api/customers/", func(w http.ResponseWriter, r *http.Request) {
		matches := pattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(matches[1])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		for _, c := range customers {
			if id == c.ID {
				data, err := json.Marshal(c)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				w.Header().Add("Content-Type", "application/json")
				w.Write(data)
				return
			}
		}
		w.WriteHeader(http.StatusNotFound)

	})

	s := http.Server{
		Addr:    ":3000",
		Handler: mux,
	}

	return &s

}
