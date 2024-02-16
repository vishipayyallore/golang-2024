package main

import (
	utl "autilities"
)

var header = utl.Header{}

func main() {
	// Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and
	// receive those values into another goroutine.
	header.DisplayHeader("Showing Channel Buffering")

	// Here we make a channel of strings buffering up to 2 values.
	messages := make(chan string, 2)

	// Because this channel is buffered, we can send these values into the channel without a corresponding concurrent receive.
	messages <- "buffered - msg 1"
	messages <- "channel - msg 2"

	// Later we can receive these two values as usual.
	utl.PLine(<-messages)
	utl.PLine(<-messages)
}

/*
Notes:

Buffered channels accept a limited number of values without a corresponding receiver for those values.
By creating a buffered channel, we can send a value into the channel without a corresponding concurrent receive.

*/
