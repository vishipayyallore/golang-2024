package main

import (
	utl "autilities"
	"errors"
	"math"
	"testing"
)

var header = utl.Header{}

/*
 */
func main() {
	/*
	 */
	header.DisplayHeader("Showing Errors Exercise")

	utl.PLine(Sqrt(9, 0.0001))
	utl.PLine(Sqrt(-2, 0.0001))
}

// Sqrt calculates the square root of x using Newton's method
func Sqrt(x float64, tolerance float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("cannot calculate square root of a negative number")
	}

	if x == 0 {
		return 0, nil
	}

	z := 1.0
	prevZ := 0.0

	for math.Abs(z-prevZ) > tolerance {
		prevZ = z
		z -= (z*z - x) / (2 * z)
	}

	return z, nil
}

func TestSqrt(t *testing.T) {
	tests := []struct {
		input     float64
		tolerance float64
		expected  float64
		wantError bool
	}{
		{input: 9, tolerance: 0.0001, expected: 3, wantError: false},            // Perfect square
		{input: 2, tolerance: 0.0001, expected: math.Sqrt(2), wantError: false}, // Non-perfect square
		{input: 0, tolerance: 0.0001, expected: 0, wantError: false},            // Zero
		{input: -1, tolerance: 0.0001, expected: 0, wantError: true},            // Negative input
	}

	for _, test := range tests {
		result, err := Sqrt(test.input, test.tolerance)
		if (err != nil) != test.wantError {
			t.Errorf("Sqrt(%v) error = %v, wantError %v", test.input, err, test.wantError)
			continue
		}
		if !test.wantError && math.Abs(result-test.expected) > test.tolerance {
			t.Errorf("Sqrt(%v) = %v, want %v (tolerance %v)", test.input, result, test.expected, test.tolerance)
		}
	}
}
