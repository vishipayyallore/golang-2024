package main

import (
	"autilities"

	"github.com/fatih/color"
)

var header = autilities.Header{}

const headerChar = '*'
const headerLength = 100
const headerColor = color.FgHiYellow
const titleColor = color.FgHiGreen
const subHeaderChar = '-'

func main() {

	header.DisplayHeader(headerChar, "Showing For Loops", headerLength, headerColor, titleColor)

	c := color.New(color.FgHiCyan)

	header.DisplayHeader(subHeaderChar, "For Loops - Only Condition", headerLength, headerColor, titleColor)
	i := 1
	for i <= 3 {

		c.Println(i)
		i = i + 1
	}

	header.DisplayHeader(subHeaderChar, "For Loops - 3 Components", headerLength, headerColor, titleColor)
	for j := 7; j <= 9; j++ {
		c.Println(j)
	}

	header.DisplayHeader(subHeaderChar, "For Loops - WITHOUT 3 Components, with break", headerLength, headerColor, titleColor)
	for {
		c.Println("loop")
		break
	}

	header.DisplayHeader(subHeaderChar, "For Loops - 3 Components, with continue", headerLength, headerColor, titleColor)
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		c.Println(n)
	}
}

/*
Notes:

1. for is Go's only looping construct. Here are three basic types of for loops.
2. The most basic type, with a single condition.
3. A classic initial/condition/after for loop.
4. for without a condition will loop repeatedly until you break out of the loop or return from the enclosing function.
5. You can also continue to the next iteration of the loop.
*/
