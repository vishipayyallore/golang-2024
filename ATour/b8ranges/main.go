package main

import (
	utl "autilities"
)

var header = utl.Header{}

func main() {

	header.DisplayHeader("Showing Range")

	showRangeOnSlices()

	showRangeOnMaps()

	showRangeOnStrings()

	showRangeAsForLoop()

	showRangeWithPow()
}

func showRangeOnSlices() {
	utl.PLine("Range on slices")

	nums := []int{2, 3, 4}
	sum := 0

	for _, num := range nums {
		sum += num
	}
	utl.PLine("Sum: ", sum)

	for i, num := range nums {
		if num == 3 {
			utl.PLine("\nIndex: ", i)
		}
	}
}

func showRangeOnMaps() {
	utl.PLine("\nRange on map")
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		utl.PFmted("%s -> %s\n", k, v)
	}

	utl.PLine("\nRange on just keys")
	for k := range kvs {
		utl.PLine("key:", k)
	}
}

func showRangeOnStrings() {
	utl.PLine("\nRange on strings")

	strData := "GoLang"
	utl.PLine("\nRange on ", strData, " string !")
	for i, c := range strData {
		utl.PFmted("%d. %d -> %c\n", i, c, c)
	}
}

func showRangeAsForLoop() {
	utl.PLine("\nRange as for loop")

	pows := []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pows {
		utl.PFmted("2**%d = %d\n", i, v)
	}
}

/*
You can skip the index or value by assigning to _.
for i, _ := range pow
for _, value := range pow
*/
func showRangeWithPow() {
	utl.PLine("\nRange with Pow")

	pow := make([]int, 10)

	// If you only want the index, you can omit the second variable.
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}

	for i, value := range pow {
		utl.PFmted("2**%d = %d\n", i, value)
	}

	// You can skip the index or value by assigning to _. If you only want the index, you can omit the second variable.
	for _, value := range pow {
		utl.PFmted("%d\n", value)
	}
}

/*
Notes:

1. Slices are a key data type in Go, giving a more powerful interface to sequences than arrays.
2. Unlike arrays, slices are typed only by the elements they contain (not the number of elements). To create an empty slice with non-zero length, use the builtin make. Here we make a slice of strings of length 3 (initially zero-valued).
3. We can set and get just like with arrays.
4. len returns the length of the slice as expected.
5. In addition to these basic operations, slices support several more that make them richer than arrays. One is the builtin append, which returns a slice containing one or more new values. Note that we need to accept a return value from append as we may get a new slice value.
6. Slices can also be copy’d. Here we create an empty slice c of the same length as slc1 and copy into c from slc1.
7. Slices support a “slice” operator with the syntax slice[low:high]. For example, this gets a slice of the elements slc1[2], slc1[3], and slc1[4].
8. This slices up to (but excluding) slc1[5].
9. And this slices up from (and including) slc1[2].
10. We can declare and initialize a variable for slice in a single line as well.
11. Slices can be composed into multi-dimensional data structures. The length of the inner slices can vary, unlike with multi-dimensional arrays.
12. Note that while slices are different types than arrays, they are rendered similarly by utl.PLine.
*/
