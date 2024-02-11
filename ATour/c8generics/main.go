package main

import (
	utl "autilities"
)

var header = utl.Header{}

func main() {
	// Starting with version 1.18, Go has added support for generics, also known as type parameters.
	header.DisplayHeader("Showing Generics")

}
