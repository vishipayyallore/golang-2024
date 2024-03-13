package main

import (
	utl "autilities"
)

var header = utl.Header{}

// We alias fmt.Println to a shorter name as we’ll use it a lot below.
var p = utl.PFmted

/*
Go offers excellent support for string formatting in the printf tradition.
*/
func main() {
	/*
	 */
	header.DisplayHeader("Showing String Formatting")

}
