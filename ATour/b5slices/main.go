package main

import (
	utl "autilities"
	"slices"

	"github.com/fatih/color"
)

var header = utl.Header{}

func main() {

	// An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array.
	// In practice, slices are much more common than arrays.
	header.DisplayHeader("Showing Slices")

	color.Cyan("Prints text in cyan.")

	showSliceDemo()

	showSliceRefToArrays()

	showSliceLiterals()

	showSliceDefaults()

	showSliceLengthCapacity()

	showArrayVsSliceDemo()

	showSliceDemo1()

	showSliceDemo2()

	showSliceDemo3()
}

func showArrayVsSliceDemo() {
	// Slices are a key data type in Go, giving a more powerful interface to sequences than arrays.
	// Unlike arrays, slices are typed only by the elements they contain (not the number of elements).
	// To create an empty slice with non-zero length, use the builtin make. Here we make a slice of strings of length 3 (initially zero-valued).
	// The make function allocates a zeroed array and returns a slice that refers to that array:
	// The type []T is a slice with elements of type T.

	var arr1 [5]int
	utl.PLine("\nArray ")
	utl.ShowTypeAndValue(arr1)

	var slc []int
	utl.PLine("\nSlice ")
	utl.ShowTypeAndValue(slc)
	utl.PLine("Uninitialized Slice: ", slc, slc == nil, len(slc) == 0)
}

func showSliceDemo() {
	utl.PLine("\nShowing Slices")

	primes := [6]int{2, 3, 5, 7, 11, 13}

	// A slice is formed by specifying two indices, a low and high bound, separated by a colon:
	// a[low : high] = This selects a half-open range which includes the first element, but excludes the last one.
	var s []int = primes[1:4]
	utl.PLine("Slice: ", s)
}

func showSliceRefToArrays() {
	utl.PLine("\nShowing Slices are like references to arrays")

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	utl.PLine("Array of Names: ", names)

	a := names[0:2]
	b := names[1:3]
	utl.PLine("Slices A, and B: ", a, b)

	b[0] = "No Name"
	utl.PLine("Slices A, and B: ", a, b)
	utl.PLine("Array of Names: ", names)
}

func showSliceLiterals() {
	// A slice literal is like an array literal without the length.
	utl.PLine("\nShowing Slice Literals")
	// This is an array literal: === [3]bool{true, true, false}
	// And this creates the same array as above, then builds a slice that references it: []bool{true, true, false}
	q := []int{2, 3, 5, 7, 11, 13}
	utl.PLine("Slice Literal: ", q)

	r := []bool{true, true, false}
	utl.PLine("Slice Literal: ", r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, true},
		{5, false},
		{7, true},
		{11, true},
		{13, false},
	}
	utl.PLine("Slice Literal: ", s)
}

// When slicing, you may omit the high or low bounds to use their defaults instead.
// The default is zero for the low bound and the length of the slice for the high bound.
func showSliceDefaults() {
	utl.PLine("\nShowing Slice Defaults")

	s := []int{2, 3, 5, 7, 11, 13}

	utl.PLine("Entire Slice: ", s)

	s = s[1:4]
	utl.PLine("Slice s[1:4]: ", s)

	s = s[:2]
	utl.PLine("Slice s[:2]: ", s)

	s = s[1:]
	utl.PLine("Slice s[1:]: ", s)
}

// A slice has both a length and a capacity.
func showSliceLengthCapacity() {
	utl.PLine("\nShowing Slice Length and Capacity")

	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	utl.PFmted("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func showSliceDemo3() {
	t := []string{"g", "h", "i"}
	utl.PLine("\nDeclare and Initialize:", t)
	utl.ShowTypeAndValue(t)
	t = append(t, "j")
	utl.ShowTypeAndValue(t)
	t = append(t, "k", "l")
	utl.ShowTypeAndValue(t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		utl.PLine("t == t2")
	}

	tdSlice := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		tdSlice[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			tdSlice[i][j] = i + j
		}
	}
	utl.PLine("\n2d Slice: ", tdSlice)

	ts := make([][]int, 3, 5)
	utl.PLine("ts: ", ts)

	ss := make([]string, 0, 5)
	utl.PLine("SS: ", ss, " len: ", len(ss), " cap: ", cap(ss))

	ss = append(ss, "A", "B")
	utl.PLine("SS: ", ss, " len: ", len(ss), " cap: ", cap(ss))

	ss = append(ss, "C", "D", "E", "F", "G", "H", "I", "J")
	utl.PLine("SS: ", ss, " len: ", len(ss), " cap: ", cap(ss))
}

func showSliceDemo2() {
	var slc1 = make([]string, 3)
	utl.ShowTypeAndValue(slc1)
	utl.PLine("\nEmpty Slice: ", slc1, " Len: ", len(slc1), " Capacity: ", cap(slc1))

	slc1[0] = "a"
	slc1[1] = "b"
	slc1[2] = "c"
	utl.PLine("set:", slc1)
	utl.PLine("get:", slc1[2])

	utl.PLine("len:", len(slc1))

	slc1 = append(slc1, "d")
	slc1 = append(slc1, "e", "f")
	utl.PLine("\nAppend:", slc1)

	c := make([]string, len(slc1))
	copy(c, slc1)
	utl.PLine("\nCopy:", c)

	l := slc1[2:5]
	utl.PLine("Slice 1:", l)

	l = slc1[:5]
	utl.PLine("Slice 2:", l)

	l = slc1[2:]
	utl.PLine("Slice 3:", l)
}

func showSliceDemo1() {
	var slcx []string
	utl.PLine("\nSlice ")
	utl.ShowTypeAndValue(slcx)
	utl.PLine("Uninitialized Slice: ", slcx, slcx == nil, len(slcx) == 0)

	slcx = make([]string, 3)
	utl.PLine("\nSlice from make: ", slcx, slcx == nil, len(slcx) == 0)
	utl.ShowTypeAndValue(slcx)
	slcx[0] = "a"
	slcx[1] = "b"
	slcx[2] = "c"
	utl.PLine("Set(slcx):", slcx)
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
