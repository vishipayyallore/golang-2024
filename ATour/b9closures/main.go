package main

import (
	utl "autilities"
)

var header = utl.Header{}

func main() {

	header.DisplayHeader("Showing Closures")

	// State will be unique to that particular function.
	sequence1 := numberSequence()
	sequence2 := numberSequence()

	pos, neg := adder(), adder()

	for i := 1; i <= 5; i++ {
		utl.PLine("Iteration: ", i)
		utl.PFmted("Sequence 1: %d\n", sequence1())
		utl.PFmted("Sequence 2: %d\n", sequence2())

		utl.PFmted("Positive: %d\n", pos(i))
		utl.PFmted("Negative: %d\n\n", neg(-2*i))
	}

}

// Go supports anonymous functions, which can form closures.
// Anonymous functions are useful when you want to define a function inline without having to name it.
func numberSequence() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

/*
Go functions may be closures. A closure is a function value that references variables from outside its body. The function may access and
assign to the referenced variables; in this sense the function is "bound" to the variables.
The adder function returns a closure. Each closure is bound to its own sum variable.
*/
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

/*
Notes:

1. Go supports anonymous functions, which can form closures.
2. Anonymous functions are useful when you want to define a function inline without having to name it.
3. This function numberSequence returns another function, which we define anonymously in the body of numberSequence.
4. The returned function closes over the variable i to form a closure.
5. Stateful closures are a useful feature of Go. This function numberSequence encapsulates state i, which is updated each time we call numberSequence.
*/
