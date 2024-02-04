// handlers/FileHandlers.go

package handlers

import (
	fileHdlers "b1staticcontentfprint/handlers"
	"net/http"
	"os"
	"time"
)

func ServeContentHandler(w http.ResponseWriter, r *http.Request) {
	customersFile, err := os.Open(fileHdlers.CustomersFilePath)

	if err != nil {
		fileHdlers.HandleError(w, err)
		return
	}
	defer customersFile.Close()

	http.ServeContent(w, r, "customerdata.csv", time.Now(), customersFile)
}
