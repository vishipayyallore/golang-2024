package main

import "fmt"

func main() {
	n1, n2 := 5, 2

	fmt.Println(n1 + n2)
	fmt.Println(n1 - n2)
	fmt.Println(n1 * n2)
	fmt.Println(n1 / n2)
	fmt.Println("Modulus: ", n1%n2)

	fmt.Println(float32(n1) / float32(n2))

	fmt.Println(n1 == n2)
	fmt.Println(n1 < n2)
	fmt.Println(n1 > n2)
}
