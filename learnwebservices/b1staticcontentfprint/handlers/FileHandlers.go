// handlers/FileHandlers.go

package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	customersFilePath = "../data/customers.csv"
)

// handleError responds with an HTTP 500 Internal Server Error and logs the error.
func handleError(w http.ResponseWriter, err error) {
	fmt.Println("Error: ", err)
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

func GetCustomerDataHandlerv1(w http.ResponseWriter, req *http.Request) {
	customerFile, err := os.Open(customersFilePath)
	if err != nil {
		handleError(w, err)
		return
	}
	defer customerFile.Close()

	// ***** Method 1 *****
	data, err := io.ReadAll(customerFile)
	if err != nil {
		handleError(w, err)
		return
	}

	fmt.Fprint(w, string(data))
}

func GetCustomerDataHandlerv2(w http.ResponseWriter, req *http.Request) {
	customerFile, err := os.Open(customersFilePath)
	if err != nil {
		handleError(w, err)
		return
	}
	defer customerFile.Close()

	// ***** Method 2 *****
	io.Copy(w, customerFile)
}
