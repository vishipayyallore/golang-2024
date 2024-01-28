package main

func swap(x, y string) (string, string) {
	return y, x
}

func getNames() (firstName, lastName string) {
	firstName = "Shri"
	lastName = " Ram"

	return // naked return
}

func getFMLNames() (firstName, middleName, lastName string) {
	return "Jai ", "Shri ", "Ram"
}

/*
Notes:

1. Go's return values may be named. If so, they are treated as variables defined at the top of the function.
2. These names should be used to document the meaning of the return values.
3. A return statement without arguments returns the current values of the results. This is known as a "naked" return.
4. Naked return statements should be used only in short functions, as with the example shown here. They can harm readability in longer functions.
*/
