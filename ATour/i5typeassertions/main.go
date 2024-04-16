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
	describe(s)

	/*
		To test whether an interface value holds a specific type, a type assertion can return two values: the underlying
		value and a boolean value that reports whether the assertion succeeded.
		t, ok := i.(T)
	*/
	s, ok := i.(string)
	utl.PLine("\ns: ", s, " ok: ", ok)
	describe(s)

	f, ok := i.(float64)
	utl.PLine("\nf: ", f, " ok: ", ok)
	describe(f)

	/*
		If i does not hold a T, the statement will trigger a panic. // panic: interface conversion: interface {} is string, not float64
	*/
	// f = i.(float64) // panic
	// utl.PLine("f: ", f)

	p, ok := i.(*P)
	utl.PLine("\np: ", p, " ok: ", ok)
	describe(p)
}

type P struct {
	Age  int
	Name string
}

/*
Under the hood, interface values can be thought of as a tuple of a value and a concrete type: (value, type). An interface value holds a
value of a specific underlying concrete type. Calling a method on an interface value executes the method of the same name on its underlying type.
*/
func describe(i interface{}) {
	utl.PLine("Describing I interface")

	utl.PFmted("(%v, %T)\n", i, i)
}
