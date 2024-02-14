package handlers

import (
	ent "d1basicrouting/entities"
	"encoding/json"
	"log"
	"net/http"
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
