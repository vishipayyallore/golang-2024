package main

import (
	utl "autilities"
	"time"
)

var header = utl.Header{}

func main() {
	// We can use channels to synchronize execution across goroutines.
	header.DisplayHeader("Showing Channel Synchronization")

	done := make(chan bool, 1)

	// Start a worker goroutine, giving it the channel to notify on.
	go worker(done)

	// Block until we receive a notification from the worker on the channel.
	<-done
	// If we removed the "<- done" line from this program, the program would exit before the worker even started.
}

// This is the function we’ll run in a goroutine. The done channel will be used to notify another goroutine that this function’s work is done.
func worker(done chan bool) {
	utl.PLine("Worker - working...")
	time.Sleep(time.Second)
	utl.PLine("Worker - done")

	// Send a value to notify that we’re done.
	done <- true
}
