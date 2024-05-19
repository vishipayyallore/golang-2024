package main

import (
	utl "autilities"
	"image"
)

var header = utl.Header{}

/*
 */
func main() {
	/*
	 */
	header.DisplayHeader("Showing Images")

	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	utl.PLine(m.Bounds())
	utl.PLine(m.At(0, 0).RGBA())
}
