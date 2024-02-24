package main

import (
	utl "autilities"
)

var header = utl.Header{}

func main() {
	// We can also use for, range syntax to iterate over values received from a channel.
	header.DisplayHeader("Showing Range over Channels")

	// We’ll iterate over 2 values in the queue channel.
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"

	// This example also showed that it’s possible to close a non-empty channel but still have the remaining values be received.
	close(queue)

	// This range iterates over each element as it’s received from queue.
	// Because we closed the channel above, the iteration terminates after receiving the 2 elements.
	for elem := range queue {
		utl.PLine(elem)
	}
}
