package main

import (
	utl "autilities"
	"fmt"
)

var header = utl.Header{}

type Person struct {
	age int
}

func (p Person) describe() string {
	return fmt.Sprintf("Person with age = %v", p.age)
}

// A Employee embeds a Person. An embedding looks like a field without a name.
type Employee struct {
	Person
	name string
}

type describer interface {
	describe() string
}

func main() {

	header.DisplayHeader("Showing Struct Embedding")
	// Go supports embedding of structs and interfaces to express a more seamless composition of types.

	// When creating structs with literals, we have to initialize the embedding explicitly;
	// here the embedded type serves as the field name.
	emp := Employee{
		Person: Person{
			age: 18,
		},
		name: "No Name",
	}

	utl.PFmted("Employee = {Age: %v, Name: %v}\n", emp.age, emp.name)
	utl.PFmted("Age = %v\n", emp.Person.age)
	utl.PFmted("Employee.describe() = %v\n", emp.describe())

	// An instance of the Employee is assigned to a variable of type describer,
	// showcasing that the Employee now fulfills the describer interface.
	var emp1 describer = emp
	utl.PFmted("emp1.describe() = %v\n", emp1.describe())
}
