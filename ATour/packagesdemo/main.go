package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Packages Demo
	fmt.Println("My favorite number is", rand.Intn(10))

	// Values Demo
	fmt.Println("go" + "lang")

	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}

/*
Notes:

1. Go has various value types including strings, integers, floats, booleans, etc. Here are a few basic examples.
*/
