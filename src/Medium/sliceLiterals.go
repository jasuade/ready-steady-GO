package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13} //This is an slice literal, that is like an array but without lengh
	fmt.Println(q)                 //Because you declare the cntent literally
	r := []bool{true, true, false, false, true, false}
	fmt.Println(r)
	//j := []int this cannot be done, you can do with {} but then is an emty array so there is no j[0]

	s := []struct { //You can create array of structs
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{}, //Note in the output that this is init as {0, false}
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

}
