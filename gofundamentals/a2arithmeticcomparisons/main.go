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

	fmt.Println("Addition: ", n1+n2)
	fmt.Println("Substraction: ", n1-n2)
	fmt.Println("Multiplication: ", n1*n2)
	fmt.Println("Integer Division: ", n1/n2) // Integer Division
	fmt.Println("Float Division: ", float64(n1)/float64(n2))
	fmt.Println("Modulus: ", n1%n2) // Modulus

	n3, n4 := 5.0, 2.0
	fmt.Println(n3 + n4)
	fmt.Println(n3 - n4)
	fmt.Println(n3 * n4)
	fmt.Println(n3 / n4)
	// Not allowed -> invalid operation: operator % not defined on n3 (variable of type float64)
	// fmt.Println("Modulus: ", n3%n4)
	fmt.Println("Modulus: ", int(n3)%int(n4))

	fmt.Println(float32(n1) / float32(n2))

	fmt.Println(n1 == n2)
	fmt.Println(n1 != n2)
	fmt.Println(n1 < n2)
	fmt.Println(n1 <= n2)
	fmt.Println(n1 > n2)
	fmt.Println(n1 >= n2)
}
