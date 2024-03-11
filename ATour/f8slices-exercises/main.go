/*
Reference: https://pkg.go.dev/golang.org/x/tour/pic#pkg-overview
*/

package main

import (
	utl "autilities"
	"image"
	"image/color"
	"image/png"
	"os"

	"golang.org/x/tour/pic"
)

const (
	dx = 256 // Width of the image
	dy = 256 // Height of the image
)

func main() {

	pic.Show(Pic)

	// Generate the image using the Pic function
	imgData := Pic(dx, dy)

	// Create a grayscale image
	// img := image.NewGray(image.Rect(0, 0, dx, dy))
	img := image.NewRGBA64(image.Rect(0, 0, dx, dy))

	// Set pixel values based on the generated data
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			// img.Set(x, y, color.Gray{Y: imgData[y][x]})
			// img.Set(x, y, color.RGBA{R: 229, G: 176, B: 135, A: imgData[y][x]})
			img.Set(x, y, color.RGBA{R: imgData[y][x]})
		}
	}

	// Create a PNG file to save the image
	file, err := os.Create("output.png")
	if err != nil {
		utl.PLine("Error creating file:", err)
		return
	}
	defer file.Close()

	// Encode the image as PNG and write it to the file
	err = png.Encode(file, img)
	if err != nil {
		utl.PLine("Error encoding PNG:", err)
		return
	}
}

/*
Implement Pic. It should return a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers.
When you run the program, it will display your picture, interpreting the integers as grayscale (well, bluescale) values.

The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and x^y.

(You need to use a loop to allocate each []uint8 inside the [][]uint8.)

(Use uint8(intValue) to convert between types.)
*/
func Pic(dx, dy int) [][]uint8 {
	utl.PLine("dx : ", dx)
	utl.PLine("dy : ", dy)

	pic := make([][]uint8, dy)
	for y := range pic {
		pic[y] = make([]uint8, dx)
		for x := range pic[y] {
			pic[y][x] = uint8((x + y) / 2)
		}
	}
	return pic
}
