package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {

	addr := ":8080"
	s := http.Server{
		Addr: addr,
	}
	wMessage := "Welcome to Customer Service Web Service.\n"
	wMessage1 := "Welcome to Customer Service Web API.\n"

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, wMessage)
	}

	// http://localhost:8080
	http.HandleFunc("/", helloHandler)

	// http://localhost:8080/api
	http.HandleFunc("/api", helloHandler)

	var getUrlHandlerFunc http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.URL.String())
	}

	// http://localhost:8080/url/ or http://localhost:8080/url/something or http://localhost:8080/url/something/else
	http.HandleFunc("/url/", getUrlHandlerFunc)

	// http://localhost:8080/handler
	http.Handle("/handler", handlerInterface(wMessage1))

	fmt.Printf("Starting Web Server at http://localhost%s\n", addr)

	go func() {
		log.Fatal(s.ListenAndServe())
	}()

	fmt.Println("Server started, press <Enter> to shutdown")
	fmt.Scanln()
	s.Shutdown(context.Background())
	fmt.Println("Server stopped")
}

type handlerInterface string

func (h handlerInterface) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Visit http://localhost:8080/handler
	w.Header().Add("X-Powered-By", "Gophers The Great !!")

	http.SetCookie(w, &http.Cookie{
		Name:    "session-id",
		Value:   "12345",
		Expires: time.Now().Add(24 * time.Hour * 365),
	})

	fmt.Fprintln(w, string(h))
	fmt.Fprintln(w, r.Header)
}
