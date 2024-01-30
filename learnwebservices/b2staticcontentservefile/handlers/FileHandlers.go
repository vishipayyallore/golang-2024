// handlers/FileHandlers.go

package handlers

import (
	"net/http"
)

const (
	customersFilePath = "../data/customers.csv"
)

func ServeFileHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, customersFilePath)
}
