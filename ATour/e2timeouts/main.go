package main

import (
	utl "autilities"
)

var header = utl.Header{}

func main() {
	// Timeouts are important for programs that connect to external resources or that otherwise need to bound execution time.
	// Implementing timeouts in Go is easy and elegant thanks to channels and select.
	header.DisplayHeader("Showing Timeouts")

}
