package main

import (
	utl "autilities"
)

var header = utl.Header{}

/*
Go has several useful functions for working with directories in the file system.
*/
func main() {
	/*
	 */
	header.DisplayHeader("Showing Directories")

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
