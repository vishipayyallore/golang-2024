package main

import (
	utl "autilities"
	"math"
)

var header = utl.Header{}

// Go supports methods defined on struct types.
type Rect struct {
	width, height int
}

// Methods can be defined for either pointer or value receiver types. Here’s an example of a value receiver.
func (r Rect) perim() int {
	return 2*r.width + 2*r.height
}

// This area method has a receiver type of *rect (pointer).
func (r *Rect) area() int {
	return r.width * r.height
}

type Vertex struct {
	X, Y float64
}

type Salary float64

/*
Go does not have classes. However, you can define methods on types. A method is a function with a special receiver argument.
We can have ONLY ONE RECEIVER for a method. The receiver appears in its own argument list between the func keyword and the method name.
*/
func main() {
	header.DisplayHeader("Showing Methods")

	r := Rect{width: 10, height: 5}
	showValueReceiver(r)

	rp := &r
	showPointerReceiver(rp)

	v := Vertex{3, 4}
	utl.PLine("\nVertex: ", v)
	utl.PLine("Abs: ", v.Abs())

	utl.PLine("Vertex: ", v)
	utl.PLine("Abs: ", Abs(v))

	utl.PLine("Vertex: ", v)
	v.Scale(10)
	utl.PLine("Scaled: ", v)

	utl.PLine("Vertex: ", v)
	Scale(&v, 10)
	utl.PLine("Scaled: ", v)

	s := Salary(50000.87)
	s.showEmployeeSalary("John Doe")
}

func (s Salary) showEmployeeSalary(emp string) {
	utl.PLine("\nMethod with Salary Receiver")
	utl.PLine("Employee: ", emp)
	utl.PLine("Salary: ", s)
}

func showValueReceiver(r Rect) {
	utl.PLine("Value Receiver -> Rectangle: ", r)
	utl.PLine("Area: ", r.area())
	utl.PLine("Perimeter: ", r.perim())
}

func showPointerReceiver(rp *Rect) {
	utl.PLine("\nPointer Receiver -> Rectangle: ", rp)
	utl.PLine("Area: ", rp.area())
	utl.PLine("Perimeter: ", rp.perim())
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/*
You can declare methods with pointer receivers. This means the receiver type has the literal syntax *T for some type T. (Also, T cannot itself
be a pointer such as *int.). For example, the Scale method here is defined on *Vertex. Methods with pointer receivers can modify the value to
which the receiver points (as Scale does here). Since methods often need to modify their receiver, pointer receivers are more common than
value receivers. With a value receiver, the Scale method operates on a copy of the original Vertex value. (This is the same behavior as for
any other function argument.) The Scale method must have a pointer receiver to change the Vertex value declared in the main function.
*/
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
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
