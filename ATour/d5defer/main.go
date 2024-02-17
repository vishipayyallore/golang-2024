package main

import (
	utl "autilities"
)

var header = utl.Header{}

func main() {
	// A defer statement defers the execution of a function until the surrounding function returns.
	header.DisplayHeader("Showing Defer")

	// The deferred call's arguments are evaluated immediately,
	// but the function call is not executed until the surrounding function returns.
	defer utl.PLine("world")

	utl.PLine("hello")

	showDefer()

	defer showDeferV2()
}

func showDefer() {
	defer utl.PLine("showDefer() world !!")

	utl.PLine("showDefer() hello ")
}

func showDeferV2() {
	utl.PLine("showDeferV2() world !!")

	utl.PLine("showDeferV2() hello ")
}
