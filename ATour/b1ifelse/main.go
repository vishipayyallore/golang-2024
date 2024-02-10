package main

import (
	utl "autilities"
	"math"
)

var header = utl.Header{}

func main() {

	header.DisplayHeader("Showing If Else")

	// the expression need not be surrounded by parentheses( )
	showSimpleIfDemo()

	// If with a short statement
	showPowerOrLimitDemo()

	if 8%4 == 0 {
		utl.PLine("8 is divisible by 4")
	}

	if 7%2 == 0 || 8%2 == 0 {
		utl.PLine("either 8 or 7 are even")
	}

	// A statement can precede conditionals; any variables declared in this statement
	// are available in the current and all subsequent branches.
	if num := 9; num < 0 {
		utl.PLine(num, " is negative")
	} else if num < 10 {
		utl.PLine(num, " has 1 digit")
	} else {
		utl.PLine(num, " has multiple digits")
	}
}

func showSimpleIfDemo() {
	utl.PLine("Simple If Demo")
	if 7%2 == 0 {
		utl.PLine("7 is even")
	} else {
		utl.PLine("7 is odd")
	}
}

func power(x, n, limit float64) float64 {
	v := 0.0

	if v = math.Pow(x, n); v < limit {
		utl.PLine("\n", v, " is less than ", limit, " so returning ", v)

		return v
	}

	utl.PLine("\n", v, " is greater than ", limit, " so returning ", limit)
	return limit
}

// Variables declared inside an if short statement are also available inside any of the else blocks.
func powerV1(x, n, limit float64) float64 {

	if v := math.Pow(x, n); v < limit {
		utl.PLine("\n", v, " is less than ", limit, " so returning ", v)

		return v
	} else {
		utl.PLine(v, " is greater than ", limit, " so returning ", limit)
	}

	// can't use v here, though
	return limit
}

func showPowerOrLimitDemo() {
	utl.PLine("Power OR Limit output is power(2, 3, 10) = ", power(2, 3, 10))
	utl.PLine("Power OR Limit output is powerV1(2, 3, 10) = ", powerV1(2, 3, 10))

	utl.PLine("Power OR Limit output is power(3, 2, 10) = ", power(3, 2, 10))
	utl.PLine("Power OR Limit output is powerV1(2, 3, 10) = ", powerV1(3, 2, 10))

	utl.PLine("Power OR Limit output is power(3, 3, 20) = ", power(3, 3, 20))
	utl.PLine("Power OR Limit output is powerV1(2, 3, 10) = ", powerV1(3, 3, 20))
}

/*
Notes:

1. There is no ternary if in Go, so you’ll need to use a full if statement even for basic conditions.
2. You can have an if statement without an else.
3. A statement can precede conditionals; any variables declared in this statement are available in all branches.
*/
