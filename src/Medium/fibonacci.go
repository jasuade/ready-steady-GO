package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	previous := 0
	total := 0
	current := 1
	return func() int {
		total = current + previous
		previous = current
		current = total
		return total
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println("Iteration number --->", i)
		fmt.Println(f())
	}
}
