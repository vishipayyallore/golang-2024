package autilities

import (
	"github.com/fatih/color"
)

func PLine(a ...interface{}) {
	c := color.New(color.FgHiCyan)

	c.Println(a...)
}
