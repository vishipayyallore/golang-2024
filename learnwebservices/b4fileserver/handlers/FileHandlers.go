// handlers/FileHandlers.go

package handlers

import (
	"net/http"
)

func GetFileServerHandlerFunc(fileServerRoute, downloadsFilePath string) http.Handler {
	return http.StripPrefix(fileServerRoute, http.FileServer(http.Dir(downloadsFilePath)))
}
