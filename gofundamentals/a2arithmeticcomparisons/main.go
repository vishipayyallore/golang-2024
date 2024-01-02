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

	header.DisplayHeader(headerChar, "Showing Arthmetic Comparisions", headerLength, headerColor, titleColor)

	n1, n2 := 5, 2

	fmt.Println(n1 + n2)
	fmt.Println(n1 - n2)
	fmt.Println(n1 * n2)
	fmt.Println(n1 / n2)
	fmt.Println("Modulus: ", n1%n2)

	fmt.Println(float32(n1) / float32(n2))

	fmt.Println(n1 == n2)
	fmt.Println(n1 < n2)
	fmt.Println(n1 > n2)
}
