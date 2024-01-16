package main

import (
	utl "autilities"
	"time"
)

var header = utl.Header{}

func main() {
	header.DisplayHeader("Hello World")

	utl.PLine("Hello Go Lang World 🙌❤️👌😊")

	utl.PLine("The time is", time.Now())
}

/*
Notes:

1. Go has various value types including strings, integers, floats, booleans, etc. Here are a few basic examples.
*/
