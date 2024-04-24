package main

import (
	utl "autilities"
	"fmt"
)

var header = utl.Header{}

/*
 */
func main() {
	/*
	 */
	header.DisplayHeader("Showing Stringers")

	a := Person{"Arthur Dent", 42}
	utl.PLine("A: ", a)
	z := Person{"Zaphod Beeblebrox", 9001}
	utl.PLine("Z: ", z)
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

/*
One of the most ubiquitous interfaces is Stringer defined by the fmt package.
A Stringer is a type that can describe itself as a string. The fmt package (and many others) look for this interface to print values.
*/
// type Stringer interface {
// 	String() string
// }
