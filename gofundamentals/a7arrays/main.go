package main

import "fmt"

func main() {

	// Int Array
	var a [5]int
	fmt.Println("Array with default values: ", a, " Length: ", len(a))

	var arr [3]string
	fmt.Println("\nArray with default values: ", arr, " Length: ", len(arr))

	arr = [3]string{"Coffee", "Espresso", "Cappuccino"}
	fmt.Println("Array after initialization: ", arr)

	fmt.Println("Array's 1st Element: ", arr[1])
	arr[1] = "Chai Tea"

	fmt.Println("Array: ", arr)

	arr2 := arr
	fmt.Println("\nArray 2: ", arr2)

	arr2[2] = "Chai Latte"

	fmt.Println("Both Arrays (arr, arr2): ", arr, arr2)
}
