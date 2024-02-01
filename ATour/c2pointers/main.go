package main

import (
	utl "autilities"
)

var header = utl.Header{}

func main() {

	header.DisplayHeader("Showing Recursion")

	fibonacci = func(n int) int {
		if n == 0 {
			return 0
		}

		if n == 1 {
			return 1
		}

		return fibonacci(n-1) + fibonacci(n-2)
	}

	for i := 1; i <= 10; i++ {
		utl.PFmted("Factorial: of %d == %d\n", i, factorial(i))
		utl.PFmted("Fibonacci: of %d == %d\n", i, fibonacci(i))
		utl.PLine()
	}

}

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

var fibonacci func(n int) int

/*
Notes:

1. A function can call itself. This is known as recursion.
2. The factorial function shown here demonstrates recursion.
3. Note the function signature of the factorial function. It takes an int and returns an int.
4. The function calls itself recursively until it reaches the base case of n == 0.
5. Recursion is not the idiomatic way to write Go code, but it is useful to understand.
6. The fibonacci function is another example of recursion.
7. Note that the fibonacci function is defined as a variable.
8. This is because we need to define it before we use it.

*/
