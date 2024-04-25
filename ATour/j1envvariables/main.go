package main

import (
	utl "autilities"
	"os"
	"strings"
)

var header = utl.Header{}

/*
Environment variables are a universal mechanism for conveying configuration information to Unix programs. Let’s look at
how to set, get, and list environment variables.
*/
func main() {
	/*
	 */
	header.DisplayHeader("Showing Environment Variables")

	// To set a key/value pair, use os.Setenv. To get a value for a key, use os.Getenv. This will return an empty string if the key isn’t present in the environment.
	os.Setenv("FOO", "1")
	utl.PLine("FOO:", os.Getenv("FOO"))
	utl.PLine("BAR:", os.Getenv("BAR"))

	// Use os.Environ to list all key/value pairs in the environment. This returns a slice of strings in the form KEY=value. You can strings.SplitN them to get the key and value. Here we print all the keys.
	utl.PLine()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		utl.PLine(pair[0], " = ", pair[1])
	}
}
