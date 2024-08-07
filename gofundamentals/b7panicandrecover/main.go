package main

import (
	"fmt"
)

/*
 */
func main() {
	fmt.Println("Showing Panic and Recover in action")

	dividend, divisor := 10, 5
	fmt.Printf("%v divided by %v is %v\n", dividend, divisor, divide(dividend, divisor))

	dividend, divisor = 10, 0
	fmt.Printf("%v divided by %v is %v\n", dividend, divisor, divide(dividend, divisor))
}

func divide(dividend, divisor int) int {
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println("Panic detected! ", msg)
		} else {
			fmt.Println("No panic detected! ", msg)
		}
	}()

	return dividend / divisor
}
