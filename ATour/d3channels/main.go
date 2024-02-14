package main

import (
	utl "autilities"
)

var header = utl.Header{}

func main() {
	// Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and
	// receive those values into another goroutine.
	header.DisplayHeader("Showing Channels")

}
