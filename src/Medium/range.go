package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow { //Range over a slice returns an array of element and its index
		fmt.Printf("2^%d = %d\n", i, v) //You can iterate over the elements
	}

	for _, value := range pow { //You can skyp the index or value assigning to _
		//Or if you just want the idex you can omit the 2nd value for i := range pow{}
		fmt.Printf("This is the index %d, this is the value %d", _, value) //Check the error
	}
}
