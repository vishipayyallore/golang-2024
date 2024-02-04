// handlers/FileHandlers.go

package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	CustomersFilePath = "../data/customers.csv"
)

// HandleError responds with an HTTP 500 Internal Server Error and logs the error.
func HandleError(w http.ResponseWriter, err error) {
	fmt.Println("Error: ", err)
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

func GetCustomerDataHandlerv1(w http.ResponseWriter, req *http.Request) {
	customerFile, err := os.Open(CustomersFilePath)
	if err != nil {
		HandleError(w, err)
		return
	}
	defer customerFile.Close()

	// ***** Method 1 *****
	data, err := io.ReadAll(customerFile)
	if err != nil {
		HandleError(w, err)
		return
	}

	fmt.Fprint(w, string(data))
}

func GetCustomerDataHandlerv2(w http.ResponseWriter, req *http.Request) {
	customerFile, err := os.Open(CustomersFilePath)
	if err != nil {
		HandleError(w, err)
		return
	}
	defer customerFile.Close()

	// ***** Method 2 *****
	io.Copy(w, customerFile)
}
