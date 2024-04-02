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

func main() {

	header.DisplayHeader(headerChar, "Showing Constants", headerLength, headerColor, titleColor)

	const a = 42 // untyped constant OR Implicitly typed constant

	autilities.PrintValueAndType(a)
	const b float32 = 3 // typed constant OR Explicitly typed constant
	autilities.PrintValueAndType(b)

	const aa = a // One constant can be assigned to another constant
	autilities.PrintValueAndType(aa)

	var i int = a
	autilities.PrintValueAndType(i)
	var f32 float32 = a
	autilities.PrintValueAndType(f32)
	f32 = b
	autilities.PrintValueAndType(f32)

	// Constants can be used in expressions
	const cccc = 2 * 5
	autilities.PrintValueAndType(cccc)
	const dddd = "Hello " + "Gophers"
	autilities.PrintValueAndType(dddd)

	const c = iota
	autilities.PrintValueAndType(c)

	// Group of Constants
	const (
		bb int = 1
		cc int = 2
	)
	autilities.PrintValueAndType(bb)
	autilities.PrintValueAndType(cc)

	const (
		aaa string = "Hello"
		bbb        = "World"
		ccc        = "!!!"
		ddd        // Uninitialized constant will take the value of the previous constant
	)
	autilities.PrintValueAndType(aaa)
	autilities.PrintValueAndType(bbb)
	autilities.PrintValueAndType(ccc)
	autilities.PrintValueAndType(ddd)

	// Grouping constants
	const (
		d = 2 * 5
		e
		f = iota // iota value is 2
		g
		h = 10 * iota // iota value is 4 as it is in 5th position
	)

	autilities.PrintValueAndType(d)
	autilities.PrintValueAndType(e)
	autilities.PrintValueAndType(f)
	autilities.PrintValueAndType(g)
	autilities.PrintValueAndType(h)

	// iota is scoped to the constant block
	// iota is reset to 0 for each constant block
	const (
		a1 = iota
	)
	autilities.PrintValueAndType(a1)

}
