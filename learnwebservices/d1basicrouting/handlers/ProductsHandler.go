package handlers

import (
	ent "d1basicrouting/entities"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
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
// GET http://localhost:8081/api/products-qs?id=1
func GetAllProductByQueryStringHandler(w http.ResponseWriter, r *http.Request) {
	idRaw := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idRaw)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
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
	}

	w.WriteHeader(http.StatusNotFound)
}

// GET http://localhost:8081/api/products/1
func GetAllProductByRouteParameterHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	log.Print(" URL: ", r.URL.Path, " | Parts : ", parts, " | Length: ", len(parts))

	if len(parts) != 4 { // path: /products/1 -> [ "" "products" "1"]
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(parts[3])
	if err != nil {
		log.Print("Error: ", err, " Parts: ", parts)
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
}
