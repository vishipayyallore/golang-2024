package autilities

import "fmt"

func ShowTypeAndValue(x interface{}) {
	fmt.Printf("Type: %T Value: %v\n", x, x)
}
