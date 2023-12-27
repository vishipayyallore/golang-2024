package main

import "fmt"

var cl, python, java bool

// Shorthand for declaring and initializing a variable not available outside a function
// k := 3

func main() {

	fmt.Println("cl, python, java", cl, python, java)

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

	var cl1, python1, java1 = true, false, "no!"
	fmt.Println("cl1, python1, java1", cl1, python1, java1)
}

/*
Notes:

1. var declares 1 or more variables. A var statement can be at package or function level
2. Go will infer the type of initialized variables. If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
3. Variables declared without a corresponding initialization are zero-valued.
4. The := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "apple" in this case.
5. Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
*/
