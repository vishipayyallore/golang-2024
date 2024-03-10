// services/CustomerService.go

package services

import (
	utls "e1shoppingcartservice/utilities"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

func CreateCustomerService() *http.Server {

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
