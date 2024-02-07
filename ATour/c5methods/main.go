package main

import (
	utl "autilities"
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

// This area method has a receiver type of *rect.
func (r *rect) area() int {
	return r.width * r.height
}

func main() {
	header.DisplayHeader("Showing Methods")

	r := rect{width: 10, height: 5}
	showValueReceiver(r)

	rp := &r
	showPointerReceiver(rp)
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
