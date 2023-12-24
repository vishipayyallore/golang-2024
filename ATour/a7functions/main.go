package main

import (
	"fmt"
)

func main() {
	fmt.Println("Sum: ", add(42, 13))

	fmt.Println("Sum: ", addv1(42, 13))

	a, b := swap("Kumar", "Manish")
	fmt.Println("Multiple Return values: ", a, b)
}

/*
Notes:

1. A function can take zero or more arguments.
2. In this example, add takes two parameters of type int.
3. Notice that the type comes after the variable name.
4. When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
5. The swap function returns two strings.
6. A return statement without arguments returns the named return values. This is known as a "naked" return.
*/
