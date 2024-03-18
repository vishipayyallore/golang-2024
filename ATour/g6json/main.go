package main

import (
	utl "autilities"
	"encoding/json"
)

var header = utl.Header{}

// We’ll use these two structs to demonstrate encoding and decoding of custom types below.
type response1 struct {
	Page   int
	Fruits []string
}

// Only exported fields will be encoded/decoded in JSON. Fields must start with capital letters to be exported.
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

/*
Go offers built-in support for JSON encoding and decoding, including to and from built-in and custom data types.
*/
func main() {
	/*
	 */
	header.DisplayHeader("Showing JSON")

	showBasicDataTypesAsJson()

	showSliceAndMapsAsJson()

	showCustomTypesAsJson()
}

func showCustomTypesAsJson() {
	utl.PLine("\nShowing custom types as JSON")

	// The JSON package can automatically encode your custom data types. It will only include exported fields in the encoded output and will by default use those names as the JSON keys.
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	utl.PLine(string(res1B))

	// You can use tags on struct field declarations to customize the encoded JSON key names. Check the definition of response2 above to see an example of such tags.
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	utl.PLine(string(res2B))
}

// Here are some for slices and maps, which encode to JSON arrays and objects as you’d expect.
func showSliceAndMapsAsJson() {
	utl.PLine("\nShowing slices and maps as JSON")

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	utl.PLine(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	utl.PLine(string(mapB))
}

// We’ll look at encoding basic data types to JSON strings. Here are some examples for atomic values.
func showBasicDataTypesAsJson() {
	utl.PLine("Showing basic data types as JSON")

	bolB, _ := json.Marshal(true)
	utl.PLine(string(bolB))

	intB, _ := json.Marshal(1)
	utl.PLine(string(intB))

	fltB, _ := json.Marshal(2.34)
	utl.PLine(string(fltB))

	strB, _ := json.Marshal("gopher")
	utl.PLine(string(strB))
}
