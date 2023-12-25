package main

func add(x int, y int) int {
	return x + y
}

func addv1(x, y int) int {
	return x + y
}

// Named return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	return
}
