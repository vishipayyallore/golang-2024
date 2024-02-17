package main

import (
	utl "autilities"
)

var header = utl.Header{}

func main() {
	// When using channels as function parameters, you can specify if a channel is meant to only send or receive values.
	// This specificity increases the type-safety of the program.
	header.DisplayHeader("Showing Channel Directions")

}

// This ping function only accepts a channel for sending values. It would be a compile-time error to try to receive on this channel.
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// The pong function accepts one channel for receives (pings) and a second for sends (pongs).
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
