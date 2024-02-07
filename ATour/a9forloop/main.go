package main

import (
	utl "autilities"
)

var header = utl.Header{}

const subHeaderChar = '-'

func main() {

	header.DisplayHeader("Showing For Loops")

	header.DisplayHeader("For Loops", utl.HeaderConfig{HeaderChar: subHeaderChar})
	// The init statement will often be a short variable declaration, and the variables declared there are visible only in the scope of the for statement.
	for j := 7; j <= 9; j++ {
		utl.PLine(j)
	}

	header.DisplayHeader("For Loops - Only Condition", utl.HeaderConfig{HeaderChar: subHeaderChar})
	// The init and post statements are optional.
	i := 1
	for i <= 3 {

		utl.PLine(i)
		i = i + 1
	}

	header.DisplayHeader("For Loops - WITHOUT 3 Components, with break", utl.HeaderConfig{HeaderChar: subHeaderChar})
	for {
		utl.PLine("loop")
		break
	}

	showForEver()

	showTable(5)

	sumNumbers(10)
}

func showForEver() {
	header.DisplayHeader("For Loops - 3 Components, with continue", utl.HeaderConfig{HeaderChar: subHeaderChar})
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		utl.PLine(n)
	}
}

func showTable(value int) {
	header.DisplayHeader("For Loops - 5 Table", utl.HeaderConfig{HeaderChar: subHeaderChar})

	for n := 1; n <= 10; n++ {
		utl.PFmted("%v * %v = %v\n", value, n, (value * n))
	}
}

// For is Go's "while"
func sumNumbers(endValue int) {
	header.DisplayHeader("For Loops - Sum of Numbers", utl.HeaderConfig{HeaderChar: subHeaderChar})

	sum := 0
	for i := 1; i <= endValue; i++ {
		sum += i
		utl.PFmted("i: %v -> Sum: %v\n", i, sum)
	}
	utl.PLine("Total: ", sum)
}

/*
Notes:

1. for is Go's only looping construct. Here are three basic types of for loops.
2. The most basic type, with a single condition.
3. A classic initial/condition/after for loop.
4. for without a condition will loop repeatedly until you break out of the loop or return from the enclosing function.
5. You can also continue to the next iteration of the loop.
*/
