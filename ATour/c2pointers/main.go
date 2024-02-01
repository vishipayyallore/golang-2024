package main

import (
	utl "autilities"
)

var header = utl.Header{}

func main() {

	header.DisplayHeader("Showing Pointers")

	i := 1
	utl.PFmted("Initial value of i inside main() = %v\n", i)

	utl.PFmted("\nBefore passByValue, i inside main() = %v\n", i)
	passByValue(i)
	utl.PFmted("After passByValue, i inside main() = %v\n", i)

	utl.PFmted("\nBefore passByReference Pointer Value, i inside main() = %v\n", &i)
	utl.PFmted("Before passByReference, i inside main() = %v\n", i)
	passByReference(&i)
	utl.PFmted("After passByReference Pointer Value, i inside main() = %v\n", &i)
	utl.PFmted("After passByReference, i inside main() = %v\n", i)

}

func passByValue(ival int) {
	utl.PFmted("In passByValue (Before), ival = %v\n", ival)
	ival = 101
	utl.PFmted("In passByValue (After), ival = %v\n", ival)
}

func passByReference(iptr *int) {
	utl.PFmted("In passByReference (Before), *iptr = %v\n", *iptr)
	*iptr = 101
	utl.PFmted("In passByReference (After), *iptr = %v\n", *iptr)
}
