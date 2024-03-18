package main

import (
	utl "autilities"
	"bytes"
	"regexp"
)

var header = utl.Header{}

/*
Go offers built-in support for regular expressions.
*/
func main() {
	/*
	 */
	header.DisplayHeader("Showing Regular Expressions")

	// This tests whether a pattern matches a string.
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	utl.PLine("Matches: ", match)

	// For other regexp tasks, you’ll need to Compile a regex.
	r, _ := regexp.Compile("p([a-z]+)ch")

	// This tests whether a pattern matches a string.
	utl.PLine("MatchString: ", r.MatchString("peach"))

	// This finds the match for the regexp.
	utl.PLine("FindString: ", r.FindString("peach punch"))

	// This also finds the first match but returns the start and end indexes for the match instead of the matching text.
	utl.PLine("FindStringIndex: ", r.FindStringIndex("peach punch"))

	// The Submatch variants include information about both the whole-pattern matches and the submatches within those matches.
	utl.PLine("FindStringSubmatch: ", r.FindStringSubmatch("peach punch"))

	// The SubmatchIndex variants return information about the indexes of matches and submatches.
	utl.PLine("FindStringSubmatchIndex: ", r.FindStringSubmatchIndex("peach punch"))

	// All variants of Find are available for byte slices as well.
	utl.PLine("FindAllString: ", r.FindAllString("peach punch pinch", -1))

	// These FindAll variants are available for byte slices as well.
	utl.PLine("FindAllStringSubmatch: ", r.FindAllStringSubmatch("peach punch pinch", -1))

	utl.PLine("FindAllStringSubmatchIndex: ", r.FindAllStringSubmatchIndex("peach punch pinch", -1))

	// Providing a non-negative integer as the second argument to these functions will limit the number of matches.
	utl.PLine("FindAllString: ", r.FindAllString("peach punch pinch", 2))

	// Our examples above had string arguments and used names like MatchString. We can also provide []byte arguments and drop String from the function name.
	utl.PLine("Match: ", r.Match([]byte("peach")))

	// When creating constants with regular expressions you can use the MustCompile variation of Compile. A plain Compile won’t work for constants because it has 2 return values.
	r = regexp.MustCompile("p([a-z]+)ch")
	utl.PLine("MustCompile: ", r)

	// The regexp package can also be used to replace subsets of strings with other values.
	utl.PLine("ReplaceAllString: ", r.ReplaceAllString("a peach", "<fruit>"))

	// The Func variant allows you to transform matched text with a given function.
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	utl.PLine("ReplaceAllFunc: ", string(out))

}
