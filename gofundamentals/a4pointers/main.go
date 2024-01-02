package main

import (
	"autilities"
	"fmt"

	"github.com/fatih/color"
)

var header = autilities.Header{}

const headerChar = '*'
const headerLength = 100
const headerColor = color.FgHiYellow
const titleColor = color.FgHiGreen

func main() {
	header.DisplayHeader(headerChar, "Showing Pointers", headerLength, headerColor, titleColor)

	s := "Hello, world!"
	autilities.PrintValueAndType(s)
	p := &s

	autilities.PrintValueAndType(p)
	fmt.Println(*p)

	*p = "Hello, Gophers!"
	autilities.PrintValueAndType(p)

	fmt.Println(s, *p)

	p = new(string)
	autilities.PrintValueAndType(p)

	fmt.Println(p, *p)
}
