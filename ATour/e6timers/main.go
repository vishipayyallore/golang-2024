package main

import (
	utl "autilities"
)

var header = utl.Header{}

func main() {
	// We often want to execute Go code at some point in the future, or repeatedly at some interval.
	// Go’s built-in timer and ticker features make both of these tasks easy.
	header.DisplayHeader("Showing Timers")

}
