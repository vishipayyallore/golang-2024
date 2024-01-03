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

const (
	Pi          = 3.14
	isManager   = true
	packageName = "a16constants"
)

func main() {
	header.DisplayHeader(headerChar, "Showing Constants", headerLength, headerColor, titleColor)

	showConstants()
}

func showConstants() {
	autilities.ShowTypeAndValue(Pi)
	autilities.ShowTypeAndValue(isManager)
	autilities.ShowTypeAndValue(packageName)
}
