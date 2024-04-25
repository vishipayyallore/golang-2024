// https://pkg.go.dev/net/http@go1.22.2#ListenAndServeTLS
// https://pkg.go.dev/net/http@go1.22.2#Server
// go run "C:\Program Files\Go\src\crypto\tls\generate_cert.go" --host localhost

package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	addr := ":8080"
	tlsConfig := &tls.Config{InsecureSkipVerify: true}
	server := &http.Server{
		Addr:      addr,
		TLSConfig: tlsConfig,
	}

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Welcome to Customer Service Web Service - Secure.\n")
	}

	http.HandleFunc("/", helloHandler)

	http.HandleFunc("/api", helloHandler)

	fmt.Printf("Starting Web Server at https://localhost%s\n", addr)

	go func() {
		log.Fatal(server.ListenAndServeTLS("./certs/cert.pem", "./certs/key.pem"))
	}()

	fmt.Println("Server started. Press <Enter> to shutdown")
	fmt.Scanln()
	server.Shutdown(context.Background())
	fmt.Println("Server stopped")
}

// http.ListenAndServeTLS(addr, "./certs/cert.pem", "./certs/key.pem", nil)
