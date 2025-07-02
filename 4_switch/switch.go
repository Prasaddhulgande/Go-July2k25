package main

import (
	"fmt"
	"time"
)

func main() {

	i := 5

	switch i {
	case 1:
		fmt.Println("One")

	case 2:
		fmt.Println("TWO")

	case 5:
		fmt.Println("FIVE")
	}

	// Multiple condition switch

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend")
	default:
		fmt.Println("It's working day")
	}

	// TYPE switch
	WoAmI := func(i interface{}) { // Interface here means any types data will come , not fix

		switch t := i.(type) {
		case int:
			fmt.Println("Its an integer")
		case string:
			fmt.Println("Its a string")
		case bool:
			fmt.Println("Its a boolean")
		default:
			fmt.Println("Other", t)
		}

	}

	WoAmI(1.1)

}
