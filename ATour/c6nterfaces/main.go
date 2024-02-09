package main

import (
	utl "autilities"
	"math"
)

var header = utl.Header{}

// Interfaces are named collections of method signatures.
type Geometry interface {
	area() float64
	perim() float64
}

// implement this interface on Rectangle and circle types.
type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

// To implement an interface in Go, we just need to implement all the methods in the interface
func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (r Rectangle) perim() float64 {
	return 2*r.width + 2*r.height
}

// To implement an interface in Go, we just need to implement all the methods in the interface
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	header.DisplayHeader("Showing Interfaces")

	showRectangle()
	showCircle()
}

// If a variable has an interface type, then we can call methods that are in the named interface.
// Here’s a generic measure function taking advantage of this to work on any geometry.
func measure(g Geometry) {
	utl.PLine("Geometry: ", g)
	utl.PLine("Area: ", g.area())
	utl.PLine("Perimeter: ", g.perim())
}

// The circle and rect struct types both implement the geometry interface so we can use instances of these
// structs as arguments to measure.
func showRectangle() {
	utl.PLine("Showing Rectangles")

	r := Rectangle{width: 3, height: 4}
	measure(r)
}

func showCircle() {
	utl.PLine("\nShowing Circles")

	c := Circle{radius: 5}
	measure(c)
}
