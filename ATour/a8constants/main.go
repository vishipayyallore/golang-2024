package main

import (
	utl "autilities"
	"math"
)

var header = utl.Header{}

const isManager rune = 'Y'
const s string = "constant"

func main() {
	header.DisplayHeader("Constants Demo")

	utl.PLine("String Constant: ", s)

	utl.PLine("Is Manager (rune Constant): ", isManager)

	const n = 500000000

	const d = 3e20 / n
	utl.PLine("Float Constant: ", d)

	utl.PLine("Int 64 Constant: ", int64(d))

	utl.PLine("Float 64: ", math.Sin(n))
}

/*
Notes:

1. Constants are declared like variables, but with the const keyword.
2. Constants can be character, string, boolean, or numeric values.
3. Constants cannot be declared using the := syntax.
*/
