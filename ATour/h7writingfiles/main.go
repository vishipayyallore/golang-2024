package main

import (
	utl "autilities"
	"os"
)

var header = utl.Header{}

const filePath string = "./data/sample.txt"

/*
Writing files in Go follows similar patterns to the ones we saw earlier for reading.
*/
func main() {
	/*
	 */
	header.DisplayHeader("Showing Writing Files")

	// To start, here’s how to dump a string (or just bytes) into a file.
	d1 := []byte("hello\ngo\n")
	err := os.WriteFile(filePath, d1, 0644)
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
