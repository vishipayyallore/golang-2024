// services/ProductService.go

package services

import (
	ent "e1shoppingcartservice/entities"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

func CreateProductService() *http.Server {

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
