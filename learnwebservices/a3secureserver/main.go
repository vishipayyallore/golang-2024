// https://pkg.go.dev/net/http@go1.21.6#ListenAndServeTLS
// go run "C:\Program Files\Go\src\crypto\tls\generate_cert.go" --host localhost

package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	addr := ":8080"

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Welcome to Customer Service Web Service - Secure.\n")
	}

	http.HandleFunc("/", helloHandler)

	http.HandleFunc("/api", helloHandler)

	fmt.Printf("Starting Web Server at https://localhost%s\n", addr)
	http.ListenAndServeTLS(addr, "./certs/cert.pem", "./certs/key.pem", nil)
}
