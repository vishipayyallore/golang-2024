package main

import (
	utl "autilities"
	"errors"
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
}

func verifyAge(age int) (int, error) {
	if age < 18 {
		// return fmt.Errorf("Age is less than 18")
		return -1, errors.New("can't work with under age 18 person")
	}

	return age, nil
}
