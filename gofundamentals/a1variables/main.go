package main

import "fmt"

func main() {
	var a string
	a = "foo"

	// fmt.Println(a)
	FmtPln(a)

	var b int = 99
	// fmt.Println(b)
	FmtPln(b)

	var c = true
	// fmt.Println(c)
	FmtPln(c)

	d := 3.1415
	// fmt.Println(d)
	FmtPln(d)

	var e int = int(d)
	// fmt.Println(e)
	FmtPln(e)
}

func FmtPln(a ...interface{}) (n int, err error) {
	return fmt.Println(a...)
}
