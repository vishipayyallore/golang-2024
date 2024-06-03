package main

import (
	"fmt"
	"math/rand"
)

/*
Every Go program is made up of packages. Programs start running in package main. By convention, the package name is the same as
the last element of the import path. For instance, the "math/rand" package comprises files that begin with the statement package rand.
*/
func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}

// Reference: <https://go.dev/tour/basics/1>
