package main

import (
	utl "autilities"
)

var header = utl.Header{}

func main() {
	// Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and
	// receive those values into another goroutine.
	header.DisplayHeader("Showing Channels")

	utl.ShowExecutablePath()

	message := "Hello Channels - Ping"

	// Create a new channel with make(chan val-type). Channels are typed by the values they convey.
	messages := make(chan string)

	// Send a value into a channel using the channel <- syntax. Here we send "Hello Channels - Ping" to the messages channel we made above, from a new goroutine.
	utl.PLine("Sending message to channel : ", message)
	go func() { messages <- "Hello Channels - Ping" }()

	//The <-channel syntax receives a value from the channel. Here we’ll receive the "Hello Channels - Ping" message we sent above and print it out.
	utl.PLine("Receiving message from channel ... ")
	msg := <-messages

	utl.PLine("Received message from channel : ", msg)
}
