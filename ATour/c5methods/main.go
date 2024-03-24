package main

import (
	utl "autilities"
	"math"
)

var header = utl.Header{}

// Go supports methods defined on struct types.
type rect struct {
	width, height int
}

// Methods can be defined for either pointer or value receiver types. Here’s an example of a value receiver.
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

// This area method has a receiver type of *rect (pointer).
func (r *rect) area() int {
	return r.width * r.height
}

type Vertex struct {
	X, Y float64
}

/*
Go does not have classes. However, you can define methods on types. A method is a function with a special receiver argument.
The receiver appears in its own argument list between the func keyword and the method name.
*/
func main() {
	header.DisplayHeader("Showing Methods")

	r := rect{width: 10, height: 5}
	showValueReceiver(r)

	rp := &r
	showPointerReceiver(rp)

	v := Vertex{3, 4}
	utl.PLine("\nVertex: ", v)
}

func showValueReceiver(r rect) {
	utl.PLine("Value Receiver -> Rectangle: ", r)
	utl.PLine("Area: ", r.area())
	utl.PLine("Perimeter: ", r.perim())
}

func showPointerReceiver(rp *rect) {
	utl.PLine("\nPointer Receiver -> Rectangle: ", rp)
	utl.PLine("Area: ", rp.area())
	utl.PLine("Perimeter: ", rp.perim())
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/*
Notes:

- Methods can be defined for either pointer or value receiver types.
- Here’s an example of a value receiver.
- This area method has a receiver type of *rect.
- The method area() has a receiver type of *rect.
- Methods with value receivers take either a value or a pointer as the receiver when they are called.
- Methods with pointer receivers take either a value or a pointer as the receiver when they are called.
- The Go compiler will help you out if you try to use a value receiver when a pointer receiver is expected.

*/
