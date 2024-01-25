// handlers/FileHandlers.go

package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func GetCustomerDataHandler(w http.ResponseWriter, req *http.Request) {
	customersFilePath := "../data/customers.csv"
	customerFile, err := os.Open(customersFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer customerFile.Close()

	// ***** Method 1 *****
	data, err := io.ReadAll(customerFile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, string(data))

	// ***** Method 2 *****
	// io.Copy(w, customerFile)
}

func ServeFileHandler(w http.ResponseWriter, r *http.Request) {
	customersFilePath := "../data/customers.csv"
	http.ServeFile(w, r, customersFilePath)
}
