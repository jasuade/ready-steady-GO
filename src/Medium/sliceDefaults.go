package main

import "fmt"

func main() {
	//Slices has lengh and capacity
	//The lenght of a slice is the num of element that contains
	//The capacity is the num of elements in the underlying array, counting for the first element of the slice
	// eg if the array is {1,2,3} and the slice is [2:2] the capacity is 0

	s := []int{2, 3, 5, 7, 11, 13}
	s = s[1:4]
	printSlice(s)
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)

	// for an array
	// a := [10]int
	//The following expresion are equivalent
	// a[0:10]
	// a[:10]
	// a[0:]
	// a[:]

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)

	//And finally nil slices
	var n []int
	fmt.Println(n, len(n), cap(n))
	if s == nil {
		fmt.Println("nil!")
	}

}

func printSlice(s []int) { //You have to define as input parameter the slice
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
