package main

import (
	"autilities"
	"fmt"
)

func main() {
	s := "Hello, world!"
	autilities.PrintValueAndType(s)
	p := &s

	autilities.PrintValueAndType(p)
	fmt.Println(*p)

	*p = "Hello, Gophers!"
	autilities.PrintValueAndType(p)

	fmt.Println(s, *p)

	p = new(string)
	autilities.PrintValueAndType(p)

	fmt.Println(p, *p)
}
