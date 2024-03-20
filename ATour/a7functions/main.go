package main

import (
	utl "autilities"
	"math"
)

var header = utl.Header{}

func main() {
	header.DisplayHeader("Functions Demo")

	num1 := 43
	num2 := 13
	num3 := 1
	total := 17

	utl.PLine("Sum ", num1, " + ", num2, " = ", add(num1, num2))

	utl.PLine("Sum: ", addTwoNumbers(num1, num2))

	utl.PFmted("Sum: %d + %d + %d = %d\n", num1, num2, num3, addThreeNumbers(num1, num2, num3))

	a, b := swap("Kumar", "Manish ")
	utl.PLine("Multiple Return values: ", a, b)

	n1, n2 := split(total)
	utl.PLine("Named (Naked) Return Values: ", n1, n2)

	// utl.PLine("output: ", split(total)) // Error: multiple-value split() in single-value context
	utl.PLine(split(total)) // This is OK

	a, b = getNames()
	utl.PLine("Named Return Values: ", a, b)

	_, c := getNames()
	utl.PLine("Named Return Values: ", c)

	a, b, c = getFMLNames()
	utl.PLine("Named Return Values: ", a, b, c)

	sum := addAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	utl.PLine("Sum: ", sum)

	sum = addAll(1, 2, 3)
	utl.PLine("Sum: ", sum)

	nums := []int{1, 2, 3, 4}
	sum = addAll(nums...)
	utl.PLine("Sum: ", sum)

	showFunctionValues()
}

// Function values may be used as function arguments and return values.
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// Functions are values too. They can be passed around just like other values.
func showFunctionValues() {
	utl.PLine("\nShowing Function Values")

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	utl.PLine("Hypotenuse: ", hypot(5, 12))

	utl.PLine("Compute: ", compute(hypot))
	utl.PLine("Compute: ", compute(math.Pow))
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
