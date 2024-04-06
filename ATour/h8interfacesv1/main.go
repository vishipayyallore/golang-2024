package main

import (
	utl "autilities"
	"math"
)

var header = utl.Header{}

/*
An interface type is defined as a set of method signatures.
*/
func main() {
	/*
		A value of interface type can hold any value that implements those methods.
	*/
	header.DisplayHeader("Showing Interfaces")

	/*
		Implicit interfaces. If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.
		In Go it is common to write methods that gracefully handle being called with a nil receiver. Note that an interface value that
		holds a nil concrete value is itself non-nil
	*/
	var t *P
	var i I = t
	i.Display()
	describe(i)

	var i3 I = new(P)
	i3.Display()
	describe(i3)

	var i1 I = &P{Age: 25, Name: "John"}
	i1.Display()
	describe(i1)

	var i2 = F(math.Pi)
	i2.Display()
	describe(i2)
}

/*
A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword. Implicit
interfaces decouple the definition of an interface from its implementation, which could then appear in any package without prearrangement.
*/
type I interface {
	Display()
}

type P struct {
	Age  int
	Name string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (p *P) Display() {
	utl.PLine("\nDisplaying Person")
	if p == nil {
		utl.PLine("Person is <nil>")
		return
	}

	utl.PLine("Age: ", p.Age)
	utl.PLine("Name: ", p.Name)
}

type F float64

func (f F) Display() {
	utl.PLine("\nDisplaying Float F.Display()")
	utl.PLine(f)
}

/*
Under the hood, interface values can be thought of as a tuple of a value and a concrete type: (value, type). An interface value holds a
value of a specific underlying concrete type. Calling a method on an interface value executes the method of the same name on its underlying type.
*/
func describe(i I) {
	utl.PFmted("(%v, %T)\n", i, i)
}
