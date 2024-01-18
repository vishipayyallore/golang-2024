package main

import (
	utl "autilities"
	"math"
)

var header = utl.Header{}

const (
	s           string = "Sri Varu"
	Pi                 = 3.14
	isManager          = true
	isManagerI  rune   = 'Y'
	packageName        = "a16constants"
	choice             = iota
	bigValue           = 1 << 100
	smallValue         = bigValue >> 99
)

func main() {
	header.DisplayHeader("Constants Demo")

	showConstantsDemo1()

	showConstantsDemo2()
}

func showConstantsDemo1() {
	utl.PLine("String Constant: ", s)

	utl.PLine("Is Manager (rune Constant): ", isManagerI)

	const n = 500000000

	const d = 3e20 / n
	utl.PLine("Float Constant: ", d)

	utl.PLine("Int 64 Constant: ", int64(d))

	utl.PLine("Float 64: ", math.Sin(n))
}

func showConstantsDemo2() {
	utl.ShowTypeAndValue(Pi)
	utl.ShowTypeAndValue(isManager)
	utl.ShowTypeAndValue(packageName)
	utl.ShowTypeAndValue(choice)

	utl.ShowTypeAndValue(toInt(smallValue))
	utl.ShowTypeAndValue(toFloat(smallValue))

	utl.ShowTypeAndValue(toFloat(bigValue))
}

func toInt(x int) int {
	return x*10 + 1
}

func toFloat(x float64) float64 {
	return x * 0.1
}

/*
Notes:

1. Constants are declared like variables, but with the const keyword.
2. Constants can be character, string, boolean, or numeric values.
3. Constants cannot be declared using the := syntax.
4. Constants are declared like variables, but with the const keyword.
5. Constants can be character, string, boolean, or numeric values.
6. Constants cannot be declared using the := syntax.
7. Constants cannot be declared using the var keyword.

*/
