package main

import (
	utl "autilities"
	"encoding/json"
	"os"
)

var header = utl.Header{}

// We’ll use these two structs to demonstrate encoding and decoding of custom types below.
type Response1 struct {
	Page   int
	Fruits []string
}

// Only exported fields will be encoded/decoded in JSON. Fields must start with capital letters to be exported.
type Response2 struct {
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

	showDecodingJson()

	showStreamJsonEncodings()
}

// In the examples above we always used bytes and strings as intermediates between the data and JSON representation on standard out.
// We can also stream JSON encodings directly to os.Writers like os.Stdout or even HTTP response bodies.
func showStreamJsonEncodings() {
	utl.PLine("\nShowing stream JSON encodings")

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}

func showDecodingJson() {
	utl.PLine("\nShowing decoding JSON")

	// We’ll start by encoding these examples of atomic values to JSON strings, which are a sequence of bytes.
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	json.Unmarshal(b, &f)
	m := f.(map[string]interface{})
	utl.PLine(m)

	// Now let’s look at decoding JSON data into Go values. Here’s an example for a generic data structure.
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// We need to provide a variable where the JSON package can put the decoded data. This map[string]interface{} will hold a map of strings to arbitrary data types.
	var dat map[string]interface{}

	// Here’s the actual decoding, and a check for associated errors.
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	utl.PLine(dat)

	// In order to use the values in the decoded map, we’ll need to convert them to their appropriate type.
	// For example here we convert the value in num to the expected float64 type.
	num := dat["num"].(float64)
	utl.PLine(num)

	// Accessing nested data requires a series of conversions.
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	utl.PLine(str1)

	// We can also decode JSON into custom data types. This has the advantages of adding additional type-safety to our programs and eliminating the need for type assertions when accessing the decoded data.
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := Response2{}
	json.Unmarshal([]byte(str), &res)
	utl.PLine(res)
	utl.PLine(res.Fruits[0])
}

func showCustomTypesAsJson() {
	utl.PLine("\nShowing custom types as JSON")

	// The JSON package can automatically encode your custom data types. It will only include exported fields in the encoded output and will by default use those names as the JSON keys.
	res1D := &Response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	utl.PLine(string(res1B))

	// You can use tags on struct field declarations to customize the encoded JSON key names. Check the definition of response2 above to see an example of such tags.
	res2D := &Response2{
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
