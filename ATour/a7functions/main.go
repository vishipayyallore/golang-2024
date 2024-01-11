package main

import (
	utl "autilities"
)

var header = utl.Header{}

func main() {
	header.DisplayHeader("Functions Demo")

	utl.PLine("Sum: ", add(42, 13))

	utl.PLine("Sum: ", addv1(42, 13))

	a, b := swap("Kumar", "Manish ")
	utl.PLine("Multiple Return values: ", a, b)

	n1, n2 := split(17)
	utl.PLine("Named (Naked) Return Values: ", n1, n2)

	// utl.PLine("output: ", split(17)) // Error: multiple-value split() in single-value context
	utl.PLine(split(17))
}

/*
Notes:

1. A function can take zero or more arguments.
2. In this example, add takes two parameters of type int.
3. Notice that the type comes after the variable name.
4. When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
5. The swap function returns two strings.
6. A return statement without arguments returns the named return values. This is known as a "naked" return.
*/
