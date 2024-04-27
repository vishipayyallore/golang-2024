package main

import (
	utl "autilities"
	"strconv"
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
}

func showErrorsDemo1(value string) {
	i, err := strconv.Atoi(value)
	if err != nil {
		utl.PFmted("couldn't convert number: %v\n", err)
		return
	}
	utl.PLine("Converted integer:", i)
}
