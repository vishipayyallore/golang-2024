package main

func add(x int, y int) int {
	return x + y
}

func addv1(x, y int) int {
	return x + y
}

// Named return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	return
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
