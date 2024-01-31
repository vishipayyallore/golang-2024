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

	for i := 1; i <= 10; i++ {
		utl.PFmted("Sequence 1: %d\n", sequence1())
		utl.PFmted("Sequence 2: %d\n", sequence2())
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
Notes:

1. Go supports anonymous functions, which can form closures.
2. Anonymous functions are useful when you want to define a function inline without having to name it.
3. This function numberSequence returns another function, which we define anonymously in the body of numberSequence.
4. The returned function closes over the variable i to form a closure.
5. Stateful closures are a useful feature of Go. This function numberSequence encapsulates state i, which is updated each time we call numberSequence.
*/
