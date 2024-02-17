package main

import (
	utl "autilities"
)

var header = utl.Header{}

func main() {
	// Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
	header.DisplayHeader("Showing Stacking defers")

	stackingDefers()

	showDeferedNumbers()

	defer showDeferedNumbersV2()

	utl.PLine("main() done !!")
}

func stackingDefers() {
	defer utl.PLine("First")
	defer utl.PLine("Second")
	defer utl.PLine("Third")
}

func showDeferedNumbers() {
	utl.PLine("Showing Defered Numbers")

	for i := 0; i < 10; i++ {
		defer utl.PLine(i)
	}

	utl.PLine("Completed Defered Numbers")
}

func showDeferedNumbersV2() {
	utl.PLine("Showing Defered Numbers V2")

	for i := 0; i < 10; i++ {
		utl.PLine(i)
	}

	utl.PLine("Completed Defered Numbers V2")
}
