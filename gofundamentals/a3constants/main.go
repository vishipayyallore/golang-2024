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

	const a = 42

	autilities.PrintValueAndType(a)
	const b float32 = 3
	autilities.PrintValueAndType(b)

	var i int = a
	autilities.PrintValueAndType(i)
	var f32 float32 = a
	autilities.PrintValueAndType(f32)
	f32 = b
	autilities.PrintValueAndType(f32)

	const c = iota
	autilities.PrintValueAndType(c)

	const (
		d = 2 * 5
		e
		f = iota
		g
		h = 10 * iota
	)

	autilities.PrintValueAndType(d)
	autilities.PrintValueAndType(e)
	autilities.PrintValueAndType(f)
	autilities.PrintValueAndType(g)
	autilities.PrintValueAndType(h)
}
