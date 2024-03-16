package main

import (
	utl "autilities"

	"github.com/fatih/color"
)

var header = utl.Header{}

const subHeaderChar = '-'

type Vertex struct {
	Lat, Long float64
}

/*
A map maps keys to values. The zero value of a map is nil. A nil map has no keys, nor can keys be added.
The make function returns a map of the given type, initialized and ready for use.
*/
func main() {

	header.DisplayHeader("Showing Maps", utl.DefaultHeaderConfig())

	config := utl.HeaderConfig{TitleColor: color.FgHiMagenta}
	header.DisplayHeader("Maps - Creating a Map", config)

	showMapsDemo1()

	showMapsDemo2()

	showMapLiterals()
}

func showMapLiterals() {
	utl.PLine("\nShowing Map Literals")

	var m = map[string]Vertex{

		// "Stores": Vertex{ 37.7749, -122.4194, },
		// If the top-level type is just a type name, you can omit it from the elements of the literal.
		"Bell Labs": {
			40.68433, -74.39967,
		},
		"Google": {
			37.42202, -122.08408,
		},
	}

	utl.PFmted("M Values: %v", m)
}

func showMapsDemo2() {
	utl.PLine("\nShowing Maps - Demo 2")

	var m1 map[string]Vertex

	utl.PFmted("M1 Values: %v", m1)
	utl.PLine("\nM1 Len: ", len(m1))
	utl.PFmted("Is M1 Nil: %v", m1 == nil)

	// The zero value of a map is nil. A nil map has no keys, nor can keys be added.
	// m1["One"] = Vertex{1, 2}
	// for k := range m1 {
	// 	utl.PLine("M1 Key: ", k)
	// }

	// The make function returns a map of the given type, initialized and ready for use.
	m1 = make(map[string]Vertex)
	m1["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

	utl.PFmted("\nM1 Values: %v", m1)
	utl.PFmted("\nM1 `Bell Labs` Value: %v", m1["Bell Labs"])
	utl.PLine("\nM1 Len: ", len(m1))
}

func showMapsDemo1() {
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
}

/*
Notes:

1. A map maps keys to values.
2. The zero value of a map is nil. A nil map has no keys, nor can keys be added.
3. The make function returns a map of the given type, initialized and ready for use.
4. The := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "apple" in this case.

*/
