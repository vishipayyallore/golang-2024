package main

import (
	utl "autilities"
	"fmt"
	"strings"
)

var header = utl.Header{}

/*
 */
func main() {
	/*
	 */
	header.DisplayHeader("Showing Stringers")

	a := Person{"Arthur Dent", 42}
	utl.PLine("A: ", a)
	z := Person{"Zaphod Beeblebrox", 9001}
	utl.PLine("Z: ", z)

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

type IPAddr [4]byte

// String method for IPAddr to implement fmt.Stringer interface
func (ip IPAddr) String() string {
	// Create a slice of strings to hold each byte converted to string
	strBytes := make([]string, len(ip))

	// Iterate over each byte in the IP address
	for i, b := range ip {
		// Convert each byte to string and add to the slice
		strBytes[i] = fmt.Sprintf("%d", b)
	}

	// Join the slice of strings with "." to form the dotted quad
	return strings.Join(strBytes, ".")
}

/*
One of the most ubiquitous interfaces is Stringer defined by the fmt package.
A Stringer is a type that can describe itself as a string. The fmt package (and many others) look for this interface to print values.
*/
// type Stringer interface {
// 	String() string
// }
