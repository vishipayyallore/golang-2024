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
	utl.PFmted("len(s) i.e number of bytes = %v", len(s))

	utl.PFmted("\nRune count: %d", utf8.RuneCountInString(s))

	showBytesInHex(s)

	rangeStringData(s)

	showDecodeRuneInStringV1(s)

	showDecodeRuneInStringV2(s)
}

func showBytesInHex(s string) {
	utl.PLine("Bytes in Hex")
	for i := 0; i < len(s); i++ {
		utl.PFmted("%x %d %c | ", s[i], s[i], s[i])
	}
}

func rangeStringData(s string) {
	utl.PLine("")
	for idx, runeValue := range s {
		utl.PFmted("%#U starts at %d\n", runeValue, idx)
	}
}

func showDecodeRuneInStringV1(s string) {
	utl.PLine("Using DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		utl.PFmted("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	if r == 't' {
		utl.PLine("found tee")
	} else if r == 'ส' {
		utl.PLine("found so sua")
	}
}

func showDecodeRuneInStringV2(s string) {
	utl.PLine("Decode Rune in String")
	for i, r := range s {
		utl.PFmted("%d\t%q\t%d\n", i, r, r)
	}
}
