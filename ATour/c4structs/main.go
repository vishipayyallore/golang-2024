package main

import (
	utl "autilities"
)

var header = utl.Header{}

func main() {

	header.DisplayHeader("Showing Structs")

	utl.PLine("Creating a new person")
	p := newPerson("John")
	utl.PLine("Person: ", p)
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
