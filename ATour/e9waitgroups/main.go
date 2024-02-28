package main

import (
	utl "autilities"
)

var header = utl.Header{}

func main() {
	// To wait for multiple goroutines to finish, we can use a wait group.
	header.DisplayHeader("Showing WaitGroups")

}
