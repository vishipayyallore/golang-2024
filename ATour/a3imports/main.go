package main

// factored import statement.
import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}

/*
Notes:

1. By convention, the package name is the same as the last element of the import path. For instance, the "math/rand" package comprises files that begin with the statement package rand.
2. Note: the environment in which these programs are executed is deterministic, so each time you run the example program rand.Intn will return the same number.
*/
