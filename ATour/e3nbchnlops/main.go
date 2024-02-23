package main

import (
	utl "autilities"
)

var header = utl.Header{}

func main() {
	// Basic sends and receives on channels are blocking. However, we can use select with a default clause to implement non-blocking sends,
	// receives, and even non-blocking multi-way selects.
	header.DisplayHeader("Showing Non-Blocking Channel Operations")

	showDemo1()
}

func showDemo1() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		utl.PLine("received message", msg)
	default:
		utl.PLine("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		utl.PLine("sent message", msg)
	default:
		utl.PLine("no message sent")
	}

	select {
	case msg := <-messages:
		utl.PLine("received message", msg)
	case sig := <-signals:
		utl.PLine("received signal", sig)
	default:
		utl.PLine("no activity")
	}
}
