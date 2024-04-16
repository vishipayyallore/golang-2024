package main

import (
	"fmt"
)

func main() {
	fmt.Println("Showing different types of loops in Go")

	i := 99

	for {
		fmt.Println(i)
		i += 1
		break
	}

	i = 5
	for i < 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

}
