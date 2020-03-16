package main

import "fmt"

func add() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
func main() {
	pos, neg := add(), add() //The variables are the type function

	for i := 0; i < 10; i++ { //In each iteration the return is the sum between the previous and the index
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

}
