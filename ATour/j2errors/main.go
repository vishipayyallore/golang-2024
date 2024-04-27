package main

import (
	utl "autilities"
	"fmt"
	"strconv"
	"time"
)

var header = utl.Header{}

/*
 */
func main() {
	/*
	 */
	header.DisplayHeader("Showing Errors")

	showErrorsDemo1("42")
	showErrorsDemo1("A42")

	if err := showCustomerErrorDemo(); err != nil {
		fmt.Println("\nCustom Error: ", err)
	}
}

func showErrorsDemo1(value string) {
	i, err := strconv.Atoi(value)
	if err != nil {
		utl.PFmted("couldn't convert number: %v\n", err)
		return
	}
	utl.PLine("Converted integer:", i)
}

type CustomError struct {
	When time.Time
	What string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func showCustomerErrorDemo() error {
	return &CustomError{
		time.Now(),
		"it didn't work",
	}
}
