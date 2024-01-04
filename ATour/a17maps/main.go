package main

import (
	utils "autilities"

	"github.com/fatih/color"
)

var header = utils.Header{}

const subHeaderChar = '-'

func main() {

	header.DisplayHeader("Showing Maps", utils.DefaultHeaderConfig())

	c := color.New(color.FgHiCyan)

	config := utils.HeaderConfig{TitleColor: color.FgHiMagenta}
	header.DisplayHeader("Maps - Creating a Map", config)

	header.DisplayHeader("Maps - Creating a Map", utils.HeaderConfig{HeaderChar: subHeaderChar})
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	c.Println("map:", m)

	v1 := m["k1"]
	c.Println("v1: ", v1)

	c.Println("len:", len(m))

	delete(m, "k2")
	c.Println("map:", m)

	_, prs := m["k2"]
	c.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	c.Println("map:", n)
}

/*
Notes:

1. A map maps keys to values.
2. The zero value of a map is nil. A nil map has no keys, nor can keys be added.
3. The make function returns a map of the given type, initialized and ready for use.
4. The := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "apple" in this case.

*/
