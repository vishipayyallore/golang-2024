package main

import (
	utl "autilities"
	"time"
)

var header = utl.Header{}

func main() {
	// Timeouts are important for programs that connect to external resources or that otherwise need to bound execution time.
	// Implementing timeouts in Go is easy and elegant thanks to channels and select.
	header.DisplayHeader("Showing Timeouts")

	showTimeoutDemo1()

	showTimeoutDemo2()
}

func showTimeoutDemo1() {
	utl.PLine("Showing Timeout Demo 1")

	c1 := make(chan string, 1)

	// suppose we’re executing an external call that returns its result on a channel c1 after 2s.
	// Note that the channel is buffered, so the send in the goroutine is nonblocking.
	// This is a common pattern to prevent goroutine leaks in case the channel is never read.
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	// Select implementing a timeout. res := <-c1 awaits the result and <-time.After awaits a value to be sent after the timeout of 1s.
	// Since select proceeds with the first receive that’s ready, we’ll take the timeout case if the operation takes more than the allowed 1s.
	select {
	case res := <-c1:
		utl.PLine(res)
	case <-time.After(1 * time.Second):
		utl.PLine("timeout 1")
	}
}

func showTimeoutDemo2() {
	utl.PLine("\nShowing Timeout Demo 2")

	c2 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	// If we allow a longer timeout of 3s, then the receive from c2 will succeed and we’ll print the result.
	select {
	case res := <-c2:
		utl.PLine(res)
	case <-time.After(3 * time.Second):
		utl.PLine("timeout 2")
	}
}
