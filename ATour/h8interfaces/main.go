package main

import (
	utl "autilities"
)

var header = utl.Header{}

/*
An interface type is defined as a set of method signatures.
*/
func main() {
	/*
	 */
	header.DisplayHeader("Showing Interfaces")

}

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}
