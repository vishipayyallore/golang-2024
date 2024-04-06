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

	var a Abser
	f := FloatValue(-math.Sqrt2)
	v := Vertex{3, 4}

	utl.PLine("A: ", a)
	utl.PLine("F: ", f)
	utl.PLine("V: ", v)

	utl.PLine("FloatValue(-math.Sqrt2).Abs() = ", f.Abs())
	utl.PLine("Vertex{3, 4}.Abs() = ", v.Abs())

	a = f // a MyFloat implements Abser
	utl.PLine("A: ", a)
	utl.PLine("a.Abs() = ", a.Abs())

	// In the following line, v is a Vertex (not *Vertex) and does NOT implement Abser.
	// cannot use v (variable of type Vertex) as Abser value in assignment: Vertex does not implement Abser (method Abs has pointer receiver)
	// a = v

	a = &v // a *Vertex implements Abser
	utl.PLine("A: ", a)
	utl.PLine("a.Abs() = ", a.Abs())

	// Implicit interfaces
	var i I = P{Age: 25, Name: "John"}
	i.Display()
	describe(i)

	var i1 I = &P{Age: 25, Name: "John"}
	i1.Display()
	describe(i1)

	var i2 = F(math.Pi)
	i2.Display()
	describe(i2)
}

type Abser interface {
	Abs() float64
}

type FloatValue float64

func (f FloatValue) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

type F float64

func (f F) Display() {
	utl.PLine("\nDisplaying Float F.Display()")
	utl.PLine(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
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
func (p P) Display() {
	utl.PLine("\nDisplaying Person")
	utl.PLine("Age: ", p.Age)
	utl.PLine("Name: ", p.Name)
}

/*
Under the hood, interface values can be thought of as a tuple of a value and a concrete type: (value, type). An interface value holds a
value of a specific underlying concrete type. Calling a method on an interface value executes the method of the same name on its underlying type.
*/
func describe(i I) {
	utl.PFmted("(%v, %T)\n", i, i)
}
