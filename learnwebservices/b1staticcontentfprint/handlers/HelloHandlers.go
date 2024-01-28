// handlers/HelloHandlers.go

package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Welcome to Customer Service Web Service.\n")
}

func GetUrlHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.URL.String())
}
