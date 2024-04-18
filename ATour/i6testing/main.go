package main

import (
	utl "autilities"
)

var header = utl.Header{}

/*
The testing package provides the tools we need to write unit tests and the go test command runs tests.
*/
func main() {
	/*
		go run .
		go test -v
		go test -bench=.
		go test -bench=. -v
	*/
	header.DisplayHeader("Showing Testing and Benchmarking in GoLang")

	var res = IntMin(2, -2)
	utl.PLine("IntMin(2, -2) = ", res)
}
