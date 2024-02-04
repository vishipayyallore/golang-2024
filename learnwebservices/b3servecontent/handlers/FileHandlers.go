// handlers/FileHandlers.go

package handlers

import (
	fileHdlers "b1staticcontentfprint/handlers"
	"net/http"
	"os"
	"time"
)

// // handleError responds with an HTTP 500 Internal Server Error and logs the error.
// func handleError(w http.ResponseWriter, err error) {
// 	fmt.Println("Error: ", err)
// 	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// }

func ServeContentHandler(w http.ResponseWriter, r *http.Request) {
	customersFile, err := os.Open(fileHdlers.CustomersFilePath)

	if err != nil {
		fileHdlers.HandleError(w, err)
		return
	}
	defer customersFile.Close()

	http.ServeContent(w, r, "customerdata.csv", time.Now(), customersFile)
}
