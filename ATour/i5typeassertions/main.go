package main

import (
	utl "autilities"
)

var header = utl.Header{}

/*
A type assertion provides access to an interface value's underlying concrete value. t := i.(T)
*/
func main() {
	/*
	 */
	header.DisplayHeader("Showing Type assertions")

	var i interface{} = "hello"

	/*
		This statement asserts that the interface value i holds the concrete type T and assigns the underlying T value to
		the variable t.
	*/
	s := i.(string)
	utl.PLine("s: ", s)

	/*
		To test whether an interface value holds a specific type, a type assertion can return two values: the underlying
		value and a boolean value that reports whether the assertion succeeded.
		t, ok := i.(T)
	*/
	s, ok := i.(string)
	utl.PLine("s: ", s, " ok: ", ok)

	f, ok := i.(float64)
	utl.PLine("f: ", f, " ok: ", ok)

	/*
		If i does not hold a T, the statement will trigger a panic. // panic: interface conversion: interface {} is string, not float64
	*/
	// f = i.(float64) // panic
	// utl.PLine("f: ", f)
}
