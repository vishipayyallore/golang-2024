package main

import (
	utl "autilities"
	"fmt"
	"os"
)

var header = utl.Header{}

/*
Command-line arguments are a common way to parameterize execution of programs.
Execution: go run . ARG1 ARG2 ARG3 ...
*/
func main() {
	/*
		os.Args provides access to raw command-line arguments. Note that the first value in this slice is the path to the program,
		and os.Args[1:] holds the arguments to the program.
	*/
	header.DisplayHeader("Showing Command-Line Arguments in GoLang")

	argsWithProg := os.Args
	fmt.Println("Arguments with Program Name: ", argsWithProg)

	argsWithoutProg := os.Args[1:]
	fmt.Println("Arguments without Program Name: ", argsWithoutProg)

	// You can get individual args with normal indexing.
	if len(os.Args) >= 4 {
		arg := os.Args[3]
		fmt.Println("Third Argument: ", arg)
	}
}
