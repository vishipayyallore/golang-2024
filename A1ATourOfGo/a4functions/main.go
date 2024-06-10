package main

import "fmt"

func main() {
	fmt.Println("42 + 13 = ", add(42, 13))
}

/*
A function can take zero or more arguments. In this example, add takes two parameters of type int.
Notice that the type comes after the variable name.
*/
func add(x int, y int) int {
	return x + y
}
