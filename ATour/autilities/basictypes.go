package autilities

import (
	"github.com/fatih/color"
)

func ShowTypeAndValue(x interface{}) {
	color.Cyan("Type: %T Value: %v\n", x, x)
	// fmt.Printf("Type: %T Value: %v\n", x, x)
}
