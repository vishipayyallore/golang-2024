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

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (r Rectangle) perim() float64 {
	return 2*r.width + 2*r.height
}

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

func measure(g Geometry) {
	utl.PLine("Geometry: ", g)
	utl.PLine("Area: ", g.area())
	utl.PLine("Perimeter: ", g.perim())
}

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
