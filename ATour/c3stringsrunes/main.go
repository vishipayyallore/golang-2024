package main

import (
	utl "autilities"
	"unicode/utf8"
)

var header = utl.Header{}

func main() {

	header.DisplayHeader("Showing Strings and Runes")

	// TODO: Convert strings into array of strings and show the details of each string
	// String literals are UTF-8 encoded text
	s := "สวัสดี"
	showStringDetails(s)
	utl.PLine("\n")

	s = "��=� ⌘"
	showStringDetails(s)
	utl.PLine("\n")

	s = "世界"
	showStringDetails(s)
	utl.PLine("\n")

	s = "Hello"
	showStringDetails(s)
}

func showStringDetails(s string) {
	utl.PFmted("String Value = %v\n", s)

	// A string is a read-only slice of bytes ([]byte).
	// The length of a string is the number of bytes it contains.
	utl.PFmted("len(s) i.e number of bytes = %v\n", len(s))

	utl.PLine("Bytes in Hex")
	for i := 0; i < len(s); i++ {
		utl.PFmted("%x %d %c | ", s[i], s[i], s[i])
	}
	utl.PFmted("")

	utl.PFmted("\nRune count: %d", utf8.RuneCountInString(s))
}
