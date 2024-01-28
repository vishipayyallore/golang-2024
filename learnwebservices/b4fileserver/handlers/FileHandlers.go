// handlers/FileHandlers.go

package handlers

import (
	"net/http"
)

func GetFileServerHandlerFunc(fileServerRoute, downloadsFilePath string) http.Handler {
	return http.StripPrefix(fileServerRoute, http.FileServer(http.Dir(downloadsFilePath)))
}

// const (
// 	customersFilePath string = "../data/customers.csv"
// )

// // handleError responds with an HTTP 500 Internal Server Error and logs the error.
// func handleError(w http.ResponseWriter, err error) {
// 	fmt.Println("Error: ", err)
// 	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// }

// func openTheFile(filePath string) (*os.File, error) {
// 	return os.Open(filePath)
// }

// func GetCustomerDataHandler(w http.ResponseWriter, req *http.Request) {
// 	customerFile, err := openTheFile(customersFilePath)

// 	if err != nil {
// 		handleError(w, err)
// 		return
// 	}
// 	defer customerFile.Close()

// 	// ***** Method 1 *****
// 	data, err := io.ReadAll(customerFile)
// 	if err != nil {
// 		handleError(w, err)
// 		return
// 	}

// 	fmt.Fprint(w, string(data))

// 	// ***** Method 2 *****
// 	// io.Copy(w, customerFile)
// }

// func ServeFileHandler(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, customersFilePath)
// }

// func ServeContentHandler(w http.ResponseWriter, r *http.Request) {
// 	customersFile, err := openTheFile(customersFilePath)

// 	if err != nil {
// 		handleError(w, err)
// 		return
// 	}
// 	defer customersFile.Close()

// 	http.ServeContent(w, r, "customerdata.csv", time.Now(), customersFile)
// }
