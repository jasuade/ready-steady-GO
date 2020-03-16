package main

import "fmt"

func main() {
	//Slices are like references to arrays
	numbers := [6]int{1, 2, 3, 4, 5, 6}
	slice := numbers[1:4]
	fmt.Println(slice)
	slice = numbers[4:6]
	fmt.Println(slice)

	names := [4]string{"Jon", "Ana", "Pol", "Eustaquio"}
	a := names[0:2]
	b := names[1:3] //As slices are like pointers they are pointing the same name
	fmt.Println(a, b)
	b[0] = "xxxxxx"   // When you change the value of what a pointer was pointing
	fmt.Println(a, b) //The result is refected in both pointers
	fmt.Println(names)

}
