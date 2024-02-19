package main

import (
	utl "autilities"
)

var header = utl.Header{}

func main() {
	// A pointer holds the memory address of a value. The type *T is a pointer to a T value. Its zero value is nil.
	header.DisplayHeader("Showing Pointers")

	showDemo1()

	showDemo2()
}

func showDemo1() {
	i, j := 42, 2701

	utl.PLine("showDemo1()")

	utl.PFmted("i: %v, j: %v\n", i, j)
	utl.PFmted("i: %v, j: %v\n", &i, &j)

	p := &i // point to i
	utl.PFmted("p: %v, p: %v\n", p, *p)
	// utl.PFmted("i: %v, j: %v\n", *(&i), *(&j))
}

func showDemo2() {
	i, j := 50, 1050

	utl.PLine("\nshowDemo2()")

	p := &i                             // point to i
	utl.PFmted("p: %v, p: %v\n", p, *p) // read i through the pointer
	*p = 21                             // set i through the pointer
	utl.PFmted("i: %v\n", i)            // see the new value of i

	p = &j                   // point to j
	*p = *p / 50             // divide j through the pointer
	utl.PFmted("j: %v\n", j) // see the new value of j
	// fmt.Println(j) // see the new value of j
}
