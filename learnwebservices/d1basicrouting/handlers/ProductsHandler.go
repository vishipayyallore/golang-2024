package handlers

import (
	ent "d1basicrouting/entities"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// GET http://localhost:8081/api/products
func GetAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(ent.Products)

	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

// GET http://localhost:8081/api/products/?id=1
func GetAllProductByQueryHandler(w http.ResponseWriter, r *http.Request) {
	idRaw := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idRaw)

	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	for _, p := range ent.Products {
		if p.ID == id {
			data, err := json.Marshal(p)
			if err != nil {
				log.Print(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Header().Add("Content-Type", "application/json")
			w.Write(data)
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}
}

// GET http://localhost:8081/api/products-qs?id=1
func GetAllProductByQueryStringHandler(w http.ResponseWriter, r *http.Request) {
	idRaw := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idRaw)

	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	for _, p := range ent.Products {
		if p.ID == id {
			data, err := json.Marshal(p)
			if err != nil {
				log.Print(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Header().Add("Content-Type", "application/json")
			w.Write(data)
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}
}
