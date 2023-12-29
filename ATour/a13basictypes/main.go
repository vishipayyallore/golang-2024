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
	showTypeAndValue(isManager)
	showTypeAndValue(Salary)
	showTypeAndValue(someValue)

	// var a1 int = 10
	// var a2 int8 = 3
	// var a3 int16 = 4
	// var a4 int32 = 5
	// var a5 int64 = 6
	// fmt.Printf("Type: %T Value: %v\n", a1, a1)
	// fmt.Printf("Type: %T Value: %v\n", a2, a2)
	// fmt.Printf("Type: %T Value: %v\n", a3, a3)
	// fmt.Printf("Type: %T Value: %v\n", a4, a4)
	// fmt.Printf("Type: %T Value: %v\n", a5, a5)

	// var b1 uint = 10
	// var b2 uint8 = 3
	// var b3 uint16 = 4
	// var b4 uint32 = 5
	// var b5 uint64 = 6
	// fmt.Printf("Type: %T Value: %v\n", b1, b1)
	// fmt.Printf("Type: %T Value: %v\n", b2, b2)
	// fmt.Printf("Type: %T Value: %v\n", b3, b3)
	// fmt.Printf("Type: %T Value: %v\n", b4, b4)
	// fmt.Printf("Type: %T Value: %v\n", b5, b5)

	// var c1 float32 = 10.1
	// var c2 float64 = 3.14
	// fmt.Printf("Type: %T Value: %v\n", c1, c1)
	// fmt.Printf("Type: %T Value: %v\n", c2, c2)

	// var d1 complex64 = 10 + 1i
	// var d2 complex128 = 3.14 + 2.7i
	// fmt.Printf("Type: %T Value: %v\n", d1, d1)
	// fmt.Printf("Type: %T Value: %v\n", d2, d2)

	// var e1 byte = 10
	// var e2 rune = 3
	// fmt.Printf("Type: %T Value: %v\n", e1, e1)
	// fmt.Printf("Type: %T Value: %v\n", e2, e2)

	// var f1 string = "Hello"
	// fmt.Printf("Type: %T Value: %v\n", f1, f1)
}

func showTypeAndValue(x interface{}) {
	fmt.Printf("Type: %T Value: %v\n", x, x)
}

/*
Notes:

1. Go has the following basic types:
	bool
	string
	int  int8  int16  int32  int64
	uint uint8 uint16 uint32 uint64 uintptr
	byte // alias for uint8
	rune // alias for int32
		// represents a Unicode code point
	float32 float64
	complex64 complex128
2. The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems.
	When you need an integer value you should use int unless you have a specific reason to use a sized or unsigned integer type.
3. The zero value is:
	0 for numeric types,
	false for the boolean type, and
	"" (the empty string) for strings.
4. Variables declared without an explicit initial value are given their zero value.
5. Constants are declared like variables, but with the const keyword.
	Constants can be character, string, boolean, or numeric values.
6. Constants cannot be declared using the := syntax.
7. Numeric constants are high-precision values.
8. An untyped constant takes the type needed by its context.
9. Try printing needInt(Big) too.

*/
