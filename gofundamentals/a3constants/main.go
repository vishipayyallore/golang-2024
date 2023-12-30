package main

import "autilities"

func main() {
	const a = 42

	autilities.PrintValueAndType(a)
	const b float32 = 3
	autilities.PrintValueAndType(b)

	var i int = a
	autilities.PrintValueAndType(i)
	var f32 float32 = a
	autilities.PrintValueAndType(f32)
	f32 = b
	autilities.PrintValueAndType(f32)

	const c = iota
	autilities.PrintValueAndType(c)

	const (
		d = 2 * 5
		e
		f = iota
		g
		h = 10 * iota
	)

	autilities.PrintValueAndType(d)
	autilities.PrintValueAndType(e)
	autilities.PrintValueAndType(f)
	autilities.PrintValueAndType(g)
	autilities.PrintValueAndType(h)
}
