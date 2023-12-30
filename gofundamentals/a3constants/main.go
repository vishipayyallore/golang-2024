package main

import "fmt"

func main() {
	const a = 42
	printValueAndType(a)
	const b float32 = 3
	printValueAndType(b)

	var i int = a
	printValueAndType(i)
	var f32 float32 = a
	printValueAndType(f32)
	f32 = b
	printValueAndType(f32)

	const c = iota
	printValueAndType(c)

	const (
		d = 2 * 5
		e
		f = iota
		g
		h = 10 * iota
	)

	printValueAndType(d)
	printValueAndType(e)
	printValueAndType(f)
	printValueAndType(g)
	printValueAndType(h)
}

func printValueAndType(a interface{}) {
	fmt.Printf("Type: %T, Value: %v\n", a, a)
}
