package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string
	fmt.Println("Slice Default: ", s, "Length:", len(s), "Capacity:", cap(s))

	s = []string{"Coffee", "Espresso", "Cappuccino"}
	fmt.Println("Slice Initialized: ", s)

	fmt.Println(s[1])
	s[1] = "Chai Tea"
	fmt.Println("Slice Modified", s)

	s = append(s, "Hot Chocolate", "Hot Tea")
	fmt.Println("Slice Appended: ", s)

	s = slices.Delete(s, 1, 2)
	fmt.Println("Slice after delete: ", s)

	s2 := s
	fmt.Println("Slice 2: ", s2)

	s[2] = "Chai Latte"
	fmt.Println("Both Slices: ", s, s2)
}
