package main

import (
	utl "autilities"
	"time"
)

var header = utl.Header{}

func main() {
	header.DisplayHeader("Hello World")

	utl.PLine("Hello Go Lang World 🙌❤️👌😊")

	utl.PLine("The time is", time.Now())
}

/*
Notes:

1. Go has various value types including strings, integers, floats, booleans, etc. Here are a few basic examples.

	header.DisplayHeader("Program Title")

	header.DisplayHeader("Program Title", utils.DefaultHeaderConfig())

	config := utils.HeaderConfig{TitleColor: color.FgHiMagenta}
	header.DisplayHeader("Program Title", config)

	header.DisplayHeader("Program Title", utils.HeaderConfig{HeaderChar: subHeaderChar})
*/
