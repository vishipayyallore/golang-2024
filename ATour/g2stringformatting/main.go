package main

import (
	utl "autilities"
	"fmt"
	"os"
)

var header = utl.Header{}

// We alias fmt.Println to a shorter name as we’ll use it a lot below.
var pf = utl.PFmted

type Point struct {
	x, y int
}

/*
Go offers excellent support for string formatting in the printf tradition.
*/
func main() {
	/*
	 */
	header.DisplayHeader("Showing String Formatting")

	p := Point{1, 2}

	// Go offers several printing “verbs” designed to format general Go values. For example, this prints an instance of our point struct.
	pf("Point: %v\n", p) // {1 2}

	// If the value is a struct, the %+v variant will include the struct’s field names.
	pf("Point: %+v\n", p) // {x:1 y:2}

	// The %#v variant prints a Go syntax representation of the value, i.e. the source code snippet that would produce that value.
	pf("Point: %#v\n", p) // main.Point{x:1, y:2}

	// To print the type of a value, use %T.
	pf("Type of Point: %T\n", p) // main.Point

	// Formatting booleans is straight-forward.
	pf("Boolean: %t\n", true)

	// There are many options for formatting integers. Use %d for standard, base-10 formatting.
	pf("Integers: %d\n", 123)

	// This prints a binary representation.
	pf("Binary: %b\n", 14)

	// This prints the character corresponding to the given integer.
	pf("Character: %c\n", 33)

	// %x provides hex encoding.
	pf("Hex: %x\n", 456)

	// There are also several formatting options for floats. For basic decimal formatting use %f.
	pf("Float: %f\n", 78.9)

	// %e and %E format the float in (slightly different versions of) scientific notation.
	pf("Scientific Notation: %e\n", 123400000.0)
	pf("Scientific Notation: %E\n", 123400000.0)

	// For basic string printing use %s.
	pf("String: %s\n", "\"string\"")

	// To double-quote strings as in Go source, use %q.
	pf("Double-Quoted String: %q\n", "\"string\"")

	// As with integers seen earlier, %x renders the string in base-16, with two output characters per byte of input.
	pf("Hex: %x\n", "hex this")

	// To print a representation of a pointer, use %p.
	pf("Pointer: %p\n", &p)

	// When formatting numbers you will often want to control the width and precision of the resulting figure. To specify the width of an integer, use a number after the % in the verb. By default the result will be right-justified and padded with spaces.
	pf("Width: |%6d|%6d|\n", 12, 345)

	// You can also specify the width of printed floats, though usually you’ll also want to restrict the decimal precision at the same time with the width.precision syntax.
	pf("Width and Precision: |%6.2f|%6.2f|\n", 1.2, 3.45)

	// To left-justify, use the - flag.
	pf("Left-Justified: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// You may also want to control width when formatting strings, especially to ensure that they align in table-like output. For basic right-justified width.
	pf("Right-Justified: |%6s|%6s|\n", "foo", "b")

	// To left-justify use the - flag as with numbers.
	pf("Left-Justified: |%-6s|%-6s|\n", "foo", "b")

	// You can also use the sprintf function to format a string without printing it. This returns a formatted string without printing it anywhere.
	s := fmt.Sprintf("sprintf: %s\n", "a string")
	pf(s)

	// You can format+print to io.Writers other than os.Stdout using Fprintf.
	fmt.Fprintf(os.Stderr, "io.Writers: %s\n", "a string")
}
