package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
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
