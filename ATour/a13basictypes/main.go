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
	fmt.Println("\nShowing Basic Types:")

	showTypeAndValue(isManager)
	showTypeAndValue(Salary)
	showTypeAndValue(someValue)

	var (
		a1 int   = 10
		a2 int8  = 3
		a3 int16 = 4
		a4 int32 = 5
		a5 int64 = 6
	)
	showTypeAndValue(a1)
	showTypeAndValue(a2)
	showTypeAndValue(a3)
	showTypeAndValue(a4)
	showTypeAndValue(a5)

	var (
		b1 uint   = 10
		b2 uint8  = 3
		b3 uint16 = 4
		b4 uint32 = 5
		b5 uint64 = 6
	)
	showTypeAndValue(b1)
	showTypeAndValue(b2)
	showTypeAndValue(b3)
	showTypeAndValue(b4)
	showTypeAndValue(b5)

	var (
		c1 float32 = 10.1
		c2 float64 = 3.14
	)
	showTypeAndValue(c1)
	showTypeAndValue(c2)

	var (
		d1 complex64  = 10 + 1i
		d2 complex128 = 3.14 + 2.7i
	)
	showTypeAndValue(d1)
	showTypeAndValue(d2)

	var (
		e1 byte = 10
		e2 rune = 3
	)
	showTypeAndValue(e1)
	showTypeAndValue(e2)

	var f1 string = "Hello"
	showTypeAndValue(f1)

	showZeroValue()
}

func showTypeAndValue(x interface{}) {
	fmt.Printf("Type: %T Value: %v\n", x, x)
}

func showZeroValue() {
	var i int
	var f float64
	var b bool
	var s string

	fmt.Println("\n\nShowing Zero Values:")

	showTypeAndValue(i)
	showTypeAndValue(f)
	showTypeAndValue(b)
	showTypeAndValue(s)
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
