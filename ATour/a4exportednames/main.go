package main

import (
	utl "autilities"
	"math"
)

var header = utl.Header{}

func main() {
	header.DisplayHeader("Exported Names Demo")

	utl.PLine("Pi is exported from math package: ", math.Pi)
}

// Notes
/*
Go, a name is exported if it begins with a capital letter
When importing a package, you can refer only to its exported names
*/
