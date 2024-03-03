// handlers/ProductsHandler.go

package handlers

import (
	ent "d1basicrouting/entities"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

// Use logrus logger
var log = logrus.New()

// GET http://localhost:8081/api/products
func GetAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("GetAllProductsHandler started")

	data, err := json.Marshal(ent.Products)

	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(data)

	log.Info("GetAllProductsHandler completed")
}

// GET http://localhost:8081/api/products/?id=1
// GET http://localhost:8081/api/products-qs?id=1
func GetAllProductByQueryStringHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("GetAllProductByQueryStringHandler started")

	idRaw := r.URL.Query().Get("id")

	id, shouldReturn := getIdfromString(idRaw, w)
	if shouldReturn {
		return
	}

	getProductByID(w, id)

	log.Info("GetAllProductByQueryStringHandler completed")
}

// GET http://localhost:8081/api/products/1
func GetAllProductByRouteParameterHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("GetAllProductByRouteParameterHandler started")

	parts := strings.Split(r.URL.Path, "/")
	log.Print(" URL: ", r.URL.Path, " | Parts : ", parts, " | Length: ", len(parts))

	if len(parts) != 4 { // path: /products/1 -> [ "" "products" "1"]
		handleError(w, fmt.Errorf("invalid route: %v", parts), http.StatusBadRequest)
		return
	}
	idRaw := parts[3]

	id, shouldReturn := getIdfromString(idRaw, w)
	if shouldReturn {
		return
	}

	getProductByID(w, id)

	log.Info("GetAllProductByRouteParameterHandler completed")
}

// handleError responds with an HTTP 500 Internal Server Error and logs the error.
func handleError(w http.ResponseWriter, err error, statusCode int) {
	log.WithError(err).WithField("status_code", statusCode).Error("Request failed")

	w.WriteHeader(statusCode)
}

func getIdfromString(idRaw string, w http.ResponseWriter) (int, bool) {
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		handleError(w, err, http.StatusBadRequest)
		return 0, true
	}
	return id, false
}

func getProductByID(w http.ResponseWriter, id int) {
	for _, p := range ent.Products {
		if p.ID == id {
			data, err := json.Marshal(p)
			if err != nil {
				log.WithError(err).Error("JSON marshaling failed")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Header().Add("Content-Type", "application/json")
			w.Write(data)
			return
		}
	}

	log.WithField("product_id", id).Warn("Product not found")
	w.WriteHeader(http.StatusNotFound)
}

// log.Printf("[%s] Error: %v\n", time.Now().Format("2006-01-02 15:04:05"), err)
