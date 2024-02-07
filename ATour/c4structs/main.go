package main

import (
	utl "autilities"
)

var header = utl.Header{}

type person struct {
	name string
	age  int
}

func main() {

	header.DisplayHeader("Showing Structs")

	createPersonStruct()

	constructNewPerson()

	accessStructFields()

	anonymousStruct()
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42

	return &p
}

// This is a function to create new person structs. It returns a pointer to the newly created struct.
func constructNewPerson() {
	utl.PLine("\nCreating a new person")

	// It’s idiomatic to encapsulate new struct creation in constructor functions.
	p := newPerson("John")
	utl.PLine("Person: ", p)
}

func createPersonStruct() {
	utl.PLine("\nCreating a new person struct")

	// This creates a new struct.
	utl.PLine(person{"Fred", 20})

	// You can name the fields when initializing a struct.
	utl.PLine(person{name: "Ann", age: 40})

	// Omitted fields will be zero-valued.
	utl.PLine(person{name: "Alice"})

	// An & prefix yields a pointer to the struct.
	utl.PLine(&person{name: "Bob", age: 50})
}

func accessStructFields() {
	utl.PLine("\nAccessing struct fields")

	s := person{name: "Sean", age: 50}
	utl.PLine(s.name)

	// You can also use dots with struct pointers - the pointers are automatically dereferenced.
	sp := &s
	utl.PLine(sp.age)

	// Structs are mutable.
	sp.age = 51
	utl.PLine(sp.age)
}

// If a struct type is only used for a single value, we don’t have to give it a name.
// The value can have an anonymous struct type.
func anonymousStruct() {
	utl.PLine("\nAnonymous struct")

	// This creates a new struct.
	a := struct {
		name string
		age  int
	}{"John", 20}

	utl.PLine(a)
}