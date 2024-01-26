package main

import (
	utl "autilities"

	"github.com/fatih/color"
)

var header = utl.Header{}

const subHeaderChar = '-'

func main() {

	header.DisplayHeader("Showing Maps", utl.DefaultHeaderConfig())

	header.DisplayHeader("Maps - Creating a Map", utl.HeaderConfig{HeaderChar: subHeaderChar})
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	utl.PLine("Map: ", m)

	v1 := m["k1"]
	utl.PLine("v1 := m['k1'] is ", v1)

	v3 := m["k3"]
	utl.PLine("v3 := m['k3'] is ", v3)

	utl.PLine("Len: ", len(m))

	delete(m, "k2")
	utl.PLine("\nMap delete(): ", m)

	clear(m)
	utl.PLine("Map clear(): ", m)

	utl.PLine("\nCurrent Map Value: ", m)
	_, kPrs := m["k1"]
	utl.PLine("Key Present m['k1']: ", kPrs)

	m["k1"] = 18
	utl.PLine("Current Map Value: ", m)
	_, kPrs = m["k1"]
	utl.PLine("Key Present m['k1']: ", kPrs)

	n := map[string]int{"foo": 1, "bar": 2}
	utl.PLine("map:", n)

	config := utl.HeaderConfig{TitleColor: color.FgHiMagenta}
	header.DisplayHeader("Maps - Creating a Map", config)
}

/*
Notes:

1. A map maps keys to values.
2. The zero value of a map is nil. A nil map has no keys, nor can keys be added.
3. The make function returns a map of the given type, initialized and ready for use.
4. The := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "apple" in this case.

*/
