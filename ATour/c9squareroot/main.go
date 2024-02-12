package main

import (
	utl "autilities"
	"math"
)

var header = utl.Header{}

func Sqrt(x float64) float64 {
	z := 1.0

	// Repeat the calculation 10 times
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)

		utl.PFmted("Iteration %d: z = %v\n", i+1, z)
	}

	return z
}

func main() {
	header.DisplayHeader("Showing Square Root")

	// Test the square root function for various values of x
	for x := 1.0; x <= 5.0; x++ {
		result := Sqrt(x)

		utl.PFmted("Square root of %v: Result = %v, math.Sqrt = %v\n", x, result, math.Sqrt(x))
		utl.PLine()
	}
}
