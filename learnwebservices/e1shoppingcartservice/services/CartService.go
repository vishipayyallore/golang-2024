package services

import (
	ent "e1shoppingcartservice/entities"
	mware "e1shoppingcartservice/middlewares"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var nextID int = 1
var carts = make([]ent.Cart, 0)

func CreateShoppingCartService() *http.Server {

	mware.CartMux.HandleFunc("/api/carts", cartsHandler)

	s := http.Server{
		Addr:    ":5005",
		Handler: &mware.LoggingMiddleware{Next: mware.CartMux},
	}

	return &s
}

func cartsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		data, err := json.Marshal(carts)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write(data)
	case http.MethodPost:
		var c ent.Cart
		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&c)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		c.ID = nextID
		nextID++
		carts = append(carts, c)
		w.WriteHeader(http.StatusCreated)
		data, err := json.Marshal(c)
		if err != nil {
			log.Print(err)
			fmt.Fprint(w, "Failed to return created cart data")
			return
		}
		w.Write(data)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
