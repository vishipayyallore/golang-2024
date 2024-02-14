package main

import (
	utl "autilities"
	"time"
)

var header = utl.Header{}

func main() {
	// A goroutine is a lightweight thread of execution.
	header.DisplayHeader("Showing Goroutines")

	f("direct")

	// To invoke this function in a goroutine, use go f(s). This new goroutine will execute concurrently with the calling one.
	go f("goroutine")

	// You can also start a goroutine for an anonymous function call.
	go func(msg string) {
		utl.PLine(msg)
	}("going")

	time.Sleep(time.Second)

	utl.PLine("done")
}

func f(from string) {
	for i := 0; i < 3; i++ {
		utl.PLine(from, ":", i)
	}
}
