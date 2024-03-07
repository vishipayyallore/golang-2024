package main

import (
	utl "autilities"
	"cmp"
	"slices"
)

var header = utl.Header{}

/*
Sometimes we’ll want to sort a collection by something other than its natural order. For example, suppose we wanted to sort
strings by their length instead of alphabetically. Here’s an example of custom sorts in Go.
*/
func main() {
	/*
	 */
	header.DisplayHeader("Showing Sorting by Functions")

	fruits := []string{"peach", "banana", "kiwi"}

	// We implement a comparison function for string lengths. cmp.Compare is helpful for this.
	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	utl.PLine("Fruits before sorting: ", fruits)
	// Now we can call slices.SortFunc with this custom comparison function to sort fruits by name length.
	slices.SortFunc(fruits, lenCmp)
	utl.PLine("Fruits after sorting: ", fruits)

	// We can use the same technique to sort a slice of values that aren’t built-in types.
	people := []Person{
		{name: "Jax", age: 37},
		{name: "TJ", age: 25},
		{name: "Alex", age: 72},
	}

	utl.PLine("People before sorting: ", people)
	slices.SortFunc(people,
		func(a, b Person) int {
			return cmp.Compare(a.age, b.age)
		})
	utl.PLine("People after sorting: ", people)
}

type Person struct {
	name string
	age  int
}
