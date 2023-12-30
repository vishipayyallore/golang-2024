package autilities

import "fmt"

func PrintValueAndType(a interface{}) {
	fmt.Printf("Type: %T, Value: %v\n", a, a)
}
