package autilities

import (
	"github.com/fatih/color"
)

// any is an alias for interface{}
func PLine(a ...interface{}) {
	c := color.New(color.FgHiCyan)

	c.Println(a...)
}
