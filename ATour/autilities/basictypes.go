package autilities

import (
	"github.com/fatih/color"
)

// Go, a name is exported if it begins with a capital letter
// any is an alias for interface{}
func ShowTypeAndValue(x interface{}) {
	color.Cyan("Type: %T Value: %v\n", x, x)
}
