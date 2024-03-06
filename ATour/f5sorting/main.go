package main

import (
	utl "autilities"
	"slices"
)

var header = utl.Header{}

/*
Go’s slices package implements sorting for builtins and user-defined types. We’ll look at sorting for builtins first
*/
func main() {
	/*
	 */
	header.DisplayHeader("Showing Sorting")

	// Sorting functions are generic, and work for any ordered built-in type.
	strs := []string{"c", "a", "b"}
	utl.PLine("Strings before sorting:", strs)
	slices.Sort(strs)
	utl.PLine("Strings after sorting:", strs)

	ints := []int{7, 2, 4}
	utl.PLine("Ints before sorting:   ", ints)
	s := slices.IsSorted(ints)
	utl.PLine("Is sorted: ", s)

	slices.Sort(ints)

	utl.PLine("Ints after sorting:   ", ints)
	s = slices.IsSorted(ints)
	utl.PLine("Is sorted: ", s)
}
