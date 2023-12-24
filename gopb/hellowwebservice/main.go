package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to Simple Web Services in Go!")
	})

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/home.html")
	})

	fmt.Println("Web server started on port 3000")

	log.Fatal(http.ListenAndServe(":3000", nil))
}
