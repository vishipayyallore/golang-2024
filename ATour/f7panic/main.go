package main

import (
	utl "autilities"
)

var header = utl.Header{}

/*
A panic typically means something went unexpectedly wrong. Mostly we use it to fail fast on errors that shouldn’t occur during normal
operation, or that we aren’t prepared to handle gracefully.
*/
func main() {
	/*
	 */
	header.DisplayHeader("Showing Panic")

}
