package main

import (
	"fmt"
)

func main() {
	// The defer statement postpone the execution of a function until the surrounding functions return
	defer fmt.Println("World")
	fmt.Println("hello")

	// you can stack several functions to be defered
	fmt.Println("Starting countdown")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Done")
}
