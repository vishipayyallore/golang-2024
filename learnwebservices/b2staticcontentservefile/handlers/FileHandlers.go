// handlers/FileHandlers.go

package handlers

import (
	fileHdlers "b1staticcontentfprint/handlers"
	"net/http"
)

func ServeFileHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, fileHdlers.CustomersFilePath)
}
