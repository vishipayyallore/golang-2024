package main

import (
	utl "autilities"
	"math/rand"
)

var header = utl.Header{}

func main() {

	header.DisplayHeader("Packages Demo")

	// Packages Demo
	utl.PLine("My favorite number is", rand.Intn(10))

	header.DisplayHeader("Values Demo")
	utl.PLine("String: go" + "lang")

	utl.PLine("Integers: 1+1 =", 1+1)
	utl.PLine("Floats: 7.0/3.0 =", 7.0/3.0)

	utl.PLine("Booleans: (true && false) =", true && false)
	utl.PLine("Booleans: (true || ) =", true || false)
	utl.PLine("Booleans: (!true) =", !true)
}

/*
Notes:

1. Go has various value types including strings, integers, floats, booleans, etc. Here are a few basic examples.
*/
