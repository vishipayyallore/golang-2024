package main

import (
	utl "autilities"
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

var header = utl.Header{}

/*
 */
func main() {
	/*
	 */
	header.DisplayHeader("Showing Images")

	m := Image{Width: 256, Height: 256}
	pic.ShowImage(m)
	utl.PLine(m.Bounds())
	utl.PLine(m.At(0, 0).RGBA())
}

// Image struct will represent our custom image.
type Image struct {
	Width, Height int
}

// Bounds returns the dimensions of the image as a rectangle.
func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.Width, img.Height)
}

// ColorModel returns the color model of the image, which is RGBA in this case.
func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

// At returns the color of the pixel at the given (x, y) coordinates.
func (img Image) At(x, y int) color.Color {
	v := uint8((x + y) / 2)
	return color.RGBA{v, v, 255, 255}
}
