package main

import "fmt"

func main() {

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)
}

/*
Notes:

1. var declares 1 or more variables.
2. Go will infer the type of initialized variables.
3. Variables declared without a corresponding initialization are zero-valued.
4. The := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "apple" in this case.
*/
