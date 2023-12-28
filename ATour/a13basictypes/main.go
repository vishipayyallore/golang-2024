package main

import (
	"fmt"
	"math/cmplx"
)

var (
	isManager bool       = false
	Salary    uint64     = 1<<64 - 1
	someValue complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", isManager, isManager)
	fmt.Printf("Type: %T Value: %v\n", Salary, Salary)
	fmt.Printf("Type: %T Value: %v\n", someValue, someValue)

	var a1 int = 10
	var a2 int8 = 3
	var a3 int16 = 4
	var a4 int32 = 5
	var a5 int64 = 6
	fmt.Printf("Type: %T Value: %v\n", a1, a1)
	fmt.Printf("Type: %T Value: %v\n", a2, a2)
	fmt.Printf("Type: %T Value: %v\n", a3, a3)
	fmt.Printf("Type: %T Value: %v\n", a4, a4)
	fmt.Printf("Type: %T Value: %v\n", a5, a5)

	var b1 uint = 10
	var b2 uint8 = 3
	var b3 uint16 = 4
	var b4 uint32 = 5
	var b5 uint64 = 6
	fmt.Printf("Type: %T Value: %v\n", b1, b1)
	fmt.Printf("Type: %T Value: %v\n", b2, b2)
	fmt.Printf("Type: %T Value: %v\n", b3, b3)
	fmt.Printf("Type: %T Value: %v\n", b4, b4)
	fmt.Printf("Type: %T Value: %v\n", b5, b5)

	var c1 float32 = 10.1
	var c2 float64 = 3.14
	fmt.Printf("Type: %T Value: %v\n", c1, c1)
	fmt.Printf("Type: %T Value: %v\n", c2, c2)

	var d1 complex64 = 10 + 1i
	var d2 complex128 = 3.14 + 2.7i
	fmt.Printf("Type: %T Value: %v\n", d1, d1)
	fmt.Printf("Type: %T Value: %v\n", d2, d2)

	var e1 byte = 10
	var e2 rune = 3
	fmt.Printf("Type: %T Value: %v\n", e1, e1)
	fmt.Printf("Type: %T Value: %v\n", e2, e2)

	var f1 string = "Hello"
	fmt.Printf("Type: %T Value: %v\n", f1, f1)
}
