package main

import (
	utl "autilities"
	"time"
)

var header = utl.Header{}

func main() {

	header.DisplayHeader("Showing Switch")

	i := 2
	utl.PLine("Write ", i, " as ")
	switch i {
	case 1:
		utl.PLine("one")
	case 2:
		utl.PLine("two")
	case 3:
		utl.PLine("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		utl.PLine("It's the weekend")
	default:
		utl.PLine("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		utl.PLine("It's before noon")
	default:
		utl.PLine("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			utl.PLine("I'm a bool")
		case int:
			utl.PLine("I'm an int")
		default:
			utl.PFmted("Don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
	whatAmI(12.49)
	whatAmI(12.49 + 1i)
	whatAmI('a')
	whatAmI(1 << 2)
	whatAmI([]int{2, 3})
}

/*
Notes:

1. A switch statement is a shorter way to write a sequence of if - else statements. It runs the first case whose value is equal to the condition expression.
2. Go only runs the selected case, not all the cases that follow. In effect, the break statement that is needed at the end of each case in those languages is provided automatically in Go.
3. Switch cases evaluate cases from top to bottom, stopping when a case succeeds.
4. Switch without a condition is the same as switch true.
5. This construct can be a clean way to write long if-then-else chains.
6. Type switches compare types instead of values. You can use this to discover the type of an interface value. In this example, the variable t will have the type corresponding to its clause.
7. You can use commas to separate multiple expressions in the same case statement. We use the optional default case in this example as well.
*/
