package main

import (
	utl "autilities"
)

var header = utl.Header{}

func main() {

	// The type [n]T is an array of n values of type T.
	header.DisplayHeader("Showing Arrays")

	showArraysDemo()

	var arr1 [5]int
	utl.PLine("\nEmpty Array: ", arr1)

	utl.PLine("Setting values ...")
	arr1[4] = 100
	utl.PLine("Set: ", arr1)
	utl.PLine("Get: ", arr1[4])

	utl.PLine("Length: ", len(arr1))

	b := [5]int{1, 2, 3, 4, 5}
	utl.PLine("Declare and Assign Array: ", b)

	var tdarr [3][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			tdarr[i][j] = i + j
		}
	}
	utl.PLine("2D Array: ", tdarr)
}

func showArraysDemo() {

	utl.PLine("Showing Arrays")

	var a [2]string
	a[0] = "Hello, "
	a[1] = "Go World"
	utl.PLine("Accessing Individual Elements -> ", a[0], a[1])
	utl.PLine("Accessing entire array -> ", a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	utl.PLine("Prime Numbers -> ", primes)
}

/*
Notes:

1. Arrays are a basic building block of Go. They can be of any type, including other arrays.
2. The length of an array is part of its type, so arrays cannot be resized. This seems like a limitation, but don't worry; Go provides a convenient way of working with arrays.
3. The type [n]T is an array of n values of type T.
4. The expression var a [10]int declares a variable a as an array of ten integers.
5. An array's length is part of its type, so arrays cannot be resized. This seems limiting, but don't worry; Go provides a convenient way of working with arrays.
*/
