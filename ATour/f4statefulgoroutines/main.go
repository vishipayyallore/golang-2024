package main

import (
	utl "autilities"
)

var header = utl.Header{}

func main() {
	/*
		We used explicit locking with mutexes to synchronize access to shared state across multiple goroutines. Another option is to use
		the built-in synchronization features of goroutines and channels to achieve the same result. This channel-based approach aligns
		with Go’s ideas of sharing memory by communicating and having each piece of data owned by exactly 1 goroutine.
	*/
	header.DisplayHeader("Showing Stateful Goroutines")

}
