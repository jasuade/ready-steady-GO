package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When it's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today is Saturday")
	case today + 1:
		fmt.Println("Tomorrow is Saturday")
	case today + 2:
		fmt.Println("The day after tomorrow is Saturday")
	default:
		fmt.Println("Ufff Saturday is still so far away . . .")
	}

	//Switch without condition is the same as switch 'true'
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Godd afternoon")
	default:
		fmt.Println("Good evening")

	}
}
