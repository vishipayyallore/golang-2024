package main

import (
	utl "autilities"
)

var header = utl.Header{}

/*
The interface type that specifies zero methods is known as the empty interface: interface{}
*/
func main() {
	/*
		An empty interface may hold values of any type. (Every type implements at least zero methods.)
	*/
	header.DisplayHeader("Showing The empty interface")

	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)

	i = &P{Age: 25, Name: "John"}
	describe(i)
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
	utl.PLine("\nDescribing I interface")

	utl.PFmted("(%v, %T)\n", i, i)
}
