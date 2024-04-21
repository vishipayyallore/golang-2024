package main

import (
	utl "autilities"
	"flag"
	"fmt"
	"os"
)

var header = utl.Header{}

/*
Command-line arguments are a common way to parameterize execution of programs.
Execution: go run . ARG1 ARG2 ARG3 ...
Execution: go run . -word=opt -numb=7 -fork -svar=flag ARG1 ARG2 ARG3 ...
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

	wordPtr := flag.String("word", "foo", "a string")
	numbPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")
	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
