package main

import (
	utl "autilities"
)

var header = utl.Header{}

/*
Go functions can be written to work on multiple types using type parameters. The type parameters of a function appear between brackets,
before the function's arguments.
*/
func main() {
	/*
	 */
	header.DisplayHeader("Showing Type parameters")

	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	utl.PLine("Found at: ", Index(si, 15))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	utl.PLine("Found at: ", Index(ss, "hello"))
}

// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}

	return -1
}
