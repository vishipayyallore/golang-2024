package main

import (
	utl "autilities"
)

var header = utl.Header{}

func main() {

	header.DisplayHeader("Showing Structs")

	createPersonStruct()

	constructNewPerson()
}

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42

	return &p
}

func constructNewPerson() {
	utl.PLine("\nCreating a new person")
	p := newPerson("John")
	utl.PLine("Person: ", p)
}

func createPersonStruct() {
	utl.PLine("Creating a new person struct")
	utl.PLine(person{"Fred", 20})
	utl.PLine(person{name: "Ann", age: 40})
	utl.PLine(person{name: "Alice"})
	utl.PLine(&person{name: "Bob", age: 50})
}
