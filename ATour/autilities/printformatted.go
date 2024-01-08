package autilities

import "github.com/fatih/color"

func PFmted(format string, a ...any) {
	c := color.New(color.FgHiCyan)

	c.Printf(format, a...)
}
