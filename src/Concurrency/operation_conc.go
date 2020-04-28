package main

import "fmt"

func main() {

	chanel := make(chan int)
	numbers := []int{1, 2, 3, 4}

	var result int

	for i := range numbers {
		go mul(numbers[i], chanel)
		result += <-chanel
	}
	fmt.Println(result)
	defer close(chanel)
}

func mul(num int, chanel chan int) {
	chanel <- num * 100
}
