package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("Write ", i, " as ")
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
		fmt.Println("It´s the weekend")
	default:
		fmt.Println("It´s weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It´s before noon")
	default:
		fmt.Println("It´s after noon")
	}

	whereAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'am a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whereAmI(true)
	whereAmI(1)
	whereAmI("hey")

}
