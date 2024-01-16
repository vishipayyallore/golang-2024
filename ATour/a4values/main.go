package main

import (
	utl "autilities"
)

var header = utl.Header{}

func main() {

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

1. Strings, which can be added together with +.
2. Integers and floats.
3. Booleans, with boolean operators as you’d expect.
*/
