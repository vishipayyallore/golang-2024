package main

import (
	utl "autilities"
	"bufio"
	"os"
)

var header = utl.Header{}

const filePath string = "./data/sample.txt"
const file1Path string = "./data/sample1.txt"

/*
Writing files in Go follows similar patterns to the ones we saw earlier for reading.
*/
func main() {
	/*
	 */
	header.DisplayHeader("Showing Writing Files")

	// To start, here’s how to dump a string (or just bytes) into a file.
	d1 := []byte("hello\ngo\n")
	err := os.WriteFile(filePath, d1, 0644)
	check(err)

	// For more granular writes, open a file for writing.
	f, err := os.Create(file1Path)
	check(err)

	// It’s idiomatic to defer a Close immediately after opening a file.
	defer f.Close()

	// You can Write byte slices as you’d expect.
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	utl.PFmted("wrote %d bytes\n", n2)

	// A WriteString is also available.
	n3, err := f.WriteString("writes\n")
	check(err)
	utl.PFmted("wrote %d bytes\n", n3)

	// Issue a Sync to flush writes to stable storage.
	f.Sync()

	// bufio provides buffered writers in addition to the buffered readers we saw earlier.
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	utl.PFmted("wrote %d bytes\n", n4)

	// Use Flush to ensure all buffered operations have been applied to the underlying writer.
	w.Flush()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
