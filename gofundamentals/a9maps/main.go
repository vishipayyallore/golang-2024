package main

import (
	"fmt"
)

func main() {
	var m map[string][]string
	fmt.Println("Map:", m, "Length:", len(m))

	m = map[string][]string{
		"coffee": {"Coffee", "Espresso", "Cappuccino"},
		"tea":    {"Hot Tea", "Chai Tea", "Chai Latte"},
	}
	fmt.Println("Map:", m, "Length:", len(m))
	fmt.Println("Coffee:", m["coffee"])
	fmt.Println("Tea:", m["tea"])

	delete(m, "tea")
	fmt.Println("Map:", m, "Length:", len(m))
	fmt.Println("Coffee:", m["coffee"])
	v, ok := m["tea"]
	fmt.Println("Tea:", v, "Present:", ok)

	m["other"] = []string{"Hot Chocolate"}
	fmt.Println("Map:", m, "Length:", len(m))
	fmt.Println("Coffee:", m["coffee"])
	v, ok = m["tea"]
	fmt.Println("Tea:", v, "Present:", ok)
	fmt.Println("Other:", m["other"])

	m2 := m
	m2["coffee"] = []string{"Coffee"}
	m["tea"] = []string{"Hot Tea"}
	fmt.Println("Map:", m, "Length:", len(m))
	fmt.Println("Map2:", m2, "Length:", len(m2))

}
