package main

import (
	utl "autilities"
	"time"
)

var header = utl.Header{}

func main() {
	// Go’s select lets you wait on multiple channel operations. Combining goroutines and channels with select is a powerful feature of Go.
	header.DisplayHeader("Showing Select")

	c1 := make(chan string)
	c2 := make(chan string)

	// Each channel will receive a value after some amount of time, to simulate
	// e.g. blocking RPC operations executing in concurrent goroutines.
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "Value One at " + time.Now().String()
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "Value Two at " + time.Now().String()
	}()

	// We’ll use select to await both of these values simultaneously, printing each one as it arrives.
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			utl.PLine(time.Now().String(), " :: Received -> ", msg1)
		case msg2 := <-c2:
			utl.PLine(time.Now().String(), " :: Received -> ", msg2)
		}
	}
}
