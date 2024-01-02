package autilities

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

// IHeader defines the header interface
type IHeader interface {
	DisplayHeader(header rune, title string, length int, colorAttribute color.Attribute)
}

// Header is a struct implementing the IHeader interface
type Header struct{}

// DisplayHeader displays the header in Go with color
func (h Header) DisplayHeader(header rune, title string, length int, headerColor, titleColor color.Attribute) {
	leftPadValue := ((length - len(title)) / 2) + len(title)
	headerValue := strings.Repeat(string(header), length)

	// Set color for printing
	headerColorFunc := color.New(headerColor).SprintFunc()
	titleColorFunc := color.New(titleColor).SprintFunc()

	headerValue = headerColorFunc(headerValue)
	fmt.Printf("\n\n%s\n", headerValue)

	title = titleColorFunc(title)
	fmt.Printf("%*s\n", leftPadValue, title)

	fmt.Printf("%s\n\n", headerValue)
}
