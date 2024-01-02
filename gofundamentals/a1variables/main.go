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

	fmt.Println(a)

	var b int = 99
	fmt.Println(b)

	var c = true
	fmt.Println(c)

	d := 3.1415
	fmt.Println(d)

	var e int = int(d)
	fmt.Println(e)
}

// func FmtPln(a ...interface{}) (n int, err error) {
// 	return fmt.Println(a...)
// }
