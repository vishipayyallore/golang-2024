package main

import (
	utl "autilities"
)

var header = utl.Header{}

func main() {

	header.DisplayHeader("Showing If Else")

	if 7%2 == 0 {
		utl.PLine("7 is even")
	} else {
		utl.PLine("7 is odd")
	}

	if 8%4 == 0 {
		utl.PLine("8 is divisible by 4")
	}

	if 7%2 == 0 || 8%2 == 0 {
		utl.PLine("either 8 or 7 are even")
	}

	if num := 9; num < 0 {
		utl.PLine(num, " is negative")
	} else if num < 10 {
		utl.PLine(num, " has 1 digit")
	} else {
		utl.PLine(num, " has multiple digits")
	}
}

/*
Notes:

1. There is no ternary if in Go, so you’ll need to use a full if statement even for basic conditions.
2. You can have an if statement without an else.
3. A statement can precede conditionals; any variables declared in this statement are available in all branches.
*/
