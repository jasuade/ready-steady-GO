package main

import "fmt"

//How to create a function that has multiple outputs
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hi", "world")
	fmt.Println(a, b)
}
