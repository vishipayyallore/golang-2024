package main

import (
	utl "autilities"
)

var header = utl.Header{}

func main() {
	// Timers are for when you want to do something once in the future.
	// Tickers are for when you want to do something repeatedly at regular intervals.
	header.DisplayHeader("Showing Tickers")

}
