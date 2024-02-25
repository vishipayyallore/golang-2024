package main

import (
	utl "autilities"
	"time"
)

var header = utl.Header{}

func main() {
	// We often want to execute Go code at some point in the future, or repeatedly at some interval.
	// Go’s built-in timer and ticker features make both of these tasks easy.
	header.DisplayHeader("Showing Timers")

	// Timers represent a single event in the future. You tell the timer how long you want to wait, and
	// it provides a channel that will be notified at that time. This timer will wait 2 seconds.
	timer1 := time.NewTimer(2 * time.Second)

	// The <-timer1.C blocks on the timer’s channel C until it sends a value indicating that the timer fired.
	<-timer1.C
	utl.PLine("Timer 1 fired")
}
