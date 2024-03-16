package main

import (
	utl "autilities"
	"os"
	"text/template"
)

var header = utl.Header{}

/*
Go offers built-in support for creating dynamic content or showing customized output to the user with the text/template package.
A sibling package named html/template provides the same API but has additional security features and should be used for generating HTML.
*/
func main() {
	/*
	 */
	header.DisplayHeader("Showing Text Templates")

	// We can create a new template and parse its body from a string.
	// Templates are a mix of static text and “actions” enclosed in {{...}} that are used to dynamically insert content.
	t1 := template.New("t1")
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}

	// Alternatively, we can use the template.Must function to panic in case Parse returns an error.
	// This is especially useful for templates initialized in the global scope.
	t1 = template.Must(t1.Parse("Value: {{.}}\n"))

	// By “executing” the template we generate its text with specific values for its actions.
	// The {{.}} action is replaced by the value passed as a parameter to Execute.
	t1.Execute(os.Stdout, "Hello, World!")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	})

}
