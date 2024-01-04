package autilities

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

// HeaderConfig defines the configuration for DisplayHeader
type HeaderConfig struct {
	HeaderChar  rune
	HeaderLen   int
	HeaderColor color.Attribute
	TitleColor  color.Attribute
}

// DefaultHeaderConfig returns the default configuration for DisplayHeader
func DefaultHeaderConfig() HeaderConfig {
	return HeaderConfig{
		HeaderChar:  '*',
		HeaderLen:   100,
		HeaderColor: color.FgHiYellow,
		TitleColor:  color.FgHiGreen,
	}
}

// IHeader defines the header interface
type IHeader interface {
	// DisplayHeader(header rune, title string, length int, colorAttribute color.Attribute)

	DisplayHeader(title string, headerConfig ...HeaderConfig)
}

/*
In Go, a type implicitly satisfies an interface if it implements all the methods declared by the interface. In your case,
the Header type implicitly satisfies the IHeader interface because it has a method with the signature defined in the interface.
The Header type implements the IHeader interface because it provides a method named DisplayHeader that matches the method signature
declared in the interface. You don't explicitly specify that Header implements IHeader as you might in some other languages. In Go,
this relationship is determined implicitly based on the method signatures. As long as the methods in Header have the same names,
parameter types, and return types as the methods in the interface, the implementation is considered valid.

// Header is a struct implementing the IHeader interface
*/
type Header struct{}

// DisplayHeader displays the header in Go with color
func (h Header) DisplayHeader(title string, headerConfig ...HeaderConfig) {

	// Set default configuration if not provided
	var hConfig HeaderConfig

	switch len(headerConfig) {
	case 0:
		// No configuration provided, use default values
		hConfig = DefaultHeaderConfig()
	case 1:
		// Use provided configuration
		hConfig = headerConfig[0]
	}

	leftPadValue := ((hConfig.HeaderLen - len(title)) / 2) + len(title)
	headerValue := strings.Repeat(string(hConfig.HeaderChar), hConfig.HeaderLen)

	// Set color for printing
	headerColorFunc := color.New(hConfig.HeaderColor).SprintFunc()
	titleColorFunc := color.New(hConfig.TitleColor).SprintFunc()

	headerValue = headerColorFunc(headerValue)
	fmt.Printf("\n\n%s\n", headerValue)

	title = titleColorFunc(title)
	fmt.Printf("%*s\n", leftPadValue, title)

	fmt.Printf("%s\n\n", headerValue)
}

/*
Usage:

// Example 1: Using default colors
header.DisplayHeader(headerChar, "Showing Maps", headerLength)

// Example 2: Providing a single color for both header and title
config := HeaderConfig{TitleColor: color.FgHiMagenta}
header.DisplayHeader(subHeaderChar, "Maps - Creating a Map", headerLength, config)

// Example 3: Providing separate colors for header and title
config := HeaderConfig{HeaderColor: color.FgHiCyan, TitleColor: color.FgHiYellow}
header.DisplayHeader(subHeaderChar, "Maps - Creating a Map", headerLength, config)

*/
