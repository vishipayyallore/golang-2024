package main

import (
	utl "autilities"
	"time"
)

var header = utl.Header{}

func main() {
	// Timers are for when you want to do something once in the future.
	// Tickers are for when you want to do something repeatedly at regular intervals.
	header.DisplayHeader("Showing Tickers")

	// Tickers use a similar mechanism to timers: a channel that is sent values.
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	// Here we’ll use the select builtin on the channel to await the values as they arrive every 500ms.
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				utl.PLine("Tick at", t)
			}
		}
	}()

	// Tickers can be stopped like timers.
	time.Sleep(1600 * time.Millisecond)

	// Once a ticker is stopped it won’t receive any more values on its channel. We’ll stop ours after 1600ms.
	utl.PLine("Ticker stopping")
	ticker.Stop()
	utl.PLine("Ticker stopped")

	done <- true

	utl.PLine("Ticker stopped")
}
