package main

import (
	ent "e1shoppingcartservice/entities"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"time"
)

var nextID int = 1
var carts = make([]ent.Cart, 0)

var cartMux = http.NewServeMux()

func createShoppingCartService() *http.Server {

	cartMux.HandleFunc("/api/carts", cartsHandler)

	s := http.Server{
		Addr:    ":5005",
		Handler: &loggingMiddleware{next: cartMux},
	}

	return &s
}

type loggingMiddleware struct {
	next http.Handler
}

func (lm loggingMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if lm.next == nil {
		lm.next = cartMux
	}

	slog.Info(fmt.Sprintf("Received %v request on route: %v", r.Method, r.URL.Path))
	now := time.Now()

	lm.next.ServeHTTP(w, r)

	slog.Info(fmt.Sprintf("Response generated for %v request on route: %v. Duration: %v", r.Method, r.URL.Path, time.Since(now)))

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
