package main

import (
	utl "autilities"
	"errors"
	"fmt"
)

var header = utl.Header{}

func main() {

	header.DisplayHeader("Showing Error")

	for _, i := range []int{15, 16, 17, 18, 19, 20} {
		if r, e := verifyAge(i); e != nil {
			utl.PLine("Please try after 18 years: ", e)
		} else {
			utl.PLine("You are selected: ", r)
		}
	}

	for _, i := range []int{15, 16, 17, 18, 19, 20} {
		if r, e := verifyAgeV2(i); e != nil {
			utl.PLine("Please try after 18 years: ", e)
		} else {
			utl.PLine("You are selected: ", r)
		}
	}

	_, e := verifyAgeV2(17)
	if ae, ok := e.(*argError); ok {
		utl.PLine(ae.arg)
		utl.PLine(ae.prob)
	}

	_, e = verifyAgeV2(42)
	if ae, ok := e.(*argError); ok {
		utl.PLine(ae.arg)
		utl.PLine(ae.prob)
	}
}

// By convention, errors are the last return value and have type error, a built-in interface.
func verifyAge(age int) (int, error) {
	if age < 18 {
		// errors.New constructs a basic error value with the given error message.
		return -1, errors.New("can't work with under age 18 person")
	}

	// A nil value in the error position indicates that there was no error.
	return age, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func verifyAgeV2(age int) (int, error) {
	if age < 18 {
		// In this case we use &argError syntax to build a new struct, supplying values for the two fields arg and prob.
		return -1, &argError{age, "can't work with under age person"}
	}

	// A nil value in the error position indicates that there was no error.
	return age, nil
}
