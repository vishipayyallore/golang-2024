package main

import (
	utl "autilities"
)

var header = utl.Header{}

func main() {
	// Generics, also known as type parameters.
	header.DisplayHeader("Showing Generics")

}

func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}

	return r
}
