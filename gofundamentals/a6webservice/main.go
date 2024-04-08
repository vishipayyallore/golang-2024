package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Handler)

	fmt.Println("Server is running at port 3000. Please open your browser and type http://localhost:3000 to see the result.")
	http.ListenAndServe(":3000", nil)

}

func Handler(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("./data/menu.txt")

	io.Copy(w, f)
}
