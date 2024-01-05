package main

import (
	"autilities"

	"github.com/fatih/color"
)

var header = autilities.Header{}

const headerChar = '*'
const headerLength = 100
const headerColor = color.FgHiYellow
const titleColor = color.FgHiGreen

const (
	Pi          = 3.14
	isManager   = true
	packageName = "a16constants"
	choice      = iota

	bigValue   = 1 << 100
	smallValue = bigValue >> 99
)

func main() {
	header.DisplayHeader(headerChar, "Showing Constants", headerLength, headerColor, titleColor)

	showConstants()
}

func showConstants() {
	autilities.ShowTypeAndValue(Pi)
	autilities.ShowTypeAndValue(isManager)
	autilities.ShowTypeAndValue(packageName)
	autilities.ShowTypeAndValue(choice)

	autilities.ShowTypeAndValue(toInt(smallValue))
	autilities.ShowTypeAndValue(toFloat(smallValue))

	autilities.ShowTypeAndValue(toFloat(bigValue))
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
4. Constants cannot be declared using the var keyword.

*/
