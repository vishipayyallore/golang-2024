package main

import (
	utl "autilities"
	"os"
)

var header = utl.Header{}

/*
 */
func main() {
	/*
	 */
	header.DisplayHeader("Showing Exit")

	// defers will not be run when using os.Exit, so this fmt.Println will never be called.
	defer utl.PLine("Print from defer !")

	// Exit with status 3.
	os.Exit(3)
}
