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

	header.DisplayHeader(headerChar, "Showing Basic Variables", headerLength, headerColor, titleColor)

	var a string
	a = "foo"

	fmt.Println("String : ", a)

	var b int = 99
	fmt.Println("Integer : ", b)

	var c = true
	fmt.Println("Boolean : ", c)

	d := 3.1415
	fmt.Println("Float 64 :", d)

	// Not allowed -> cannot use d (variable of type float64) as int value in variable declaration
	// var e int = d
	var e int = int(d) // Explicit Type conversion
	fmt.Println("Integer : ", e)

	var f int8 = 127
	fmt.Println("Integer-8 : ", f)

	// Not allowed -> cannot use f (variable of type int8) as int16 value in variable declaration
	// var g int16 = f
	var g int16 = int16(f) // Explicit Type conversion
	fmt.Println("Integer-16 : ", g)
}
