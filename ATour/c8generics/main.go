package main

import (
	utl "autilities"
)

var header = utl.Header{}

func main() {
	// Generics, also known as type parameters.
	header.DisplayHeader("Showing Generics")

	var m = map[int]string{1: "2", 2: "4", 4: "8"}
	utl.PLine("MapKeys(m) = ", MapKeys(m))
}

/*
 * As an example of a generic function, MapKeys takes a map of any type and returns a slice of its keys.
 * This function has two type parameters - K and V; K has the comparable constraint, meaning that we can compare values
 * of this type with the == and != operators. This is required for map keys in Go. V has the any constraint,
 * meaning that it’s not restricted in any way (any is an alias for interface{}).
 */
func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}

	return r
}

type element[T any] struct {
	next *element[T]
	val  T
}

// As an example of a generic type, List is a singly-linked list with values of any type.
type List[T any] struct {
	head, tail *element[T]
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}
func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}
