package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}

/*
Notes:

1. Constants are declared like variables, but with the const keyword.
2. Constants can be character, string, boolean, or numeric values.
3. Constants cannot be declared using the := syntax.
*/
