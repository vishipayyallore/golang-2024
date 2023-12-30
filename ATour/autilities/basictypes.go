package autilities

import (
	"fmt"

	"github.com/fatih/color"
)

func ShowTypeAndValue(x interface{}) {
	color.Magenta("Type: %T Value: %v\n", x, x)
	fmt.Printf("Type: %T Value: %v\n", x, x)
}
