package main

import "fmt"

func main() {
	var num []int
	printSlice(num)

	num = append(num, 0) //Build-in append works on nil slices
	printSlice(num)
	num = append(num, 1)
	printSlice(num)                //the slice grows as needed
	num = append(num, 2, 3, 4, 23) //more than one elements can be added at the same time
	printSlice(num)
}

func printSlice(num []int) {
	fmt.Printf("%v, len=%d, cap=%d \n", num, len(num), cap(num))
}
