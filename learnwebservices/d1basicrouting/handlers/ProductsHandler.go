package handlers

import (
	ent "d1basicrouting/entities"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

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
