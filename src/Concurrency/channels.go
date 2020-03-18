package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum //Sends the sum to the channel c
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c //Receive the results into x and y variables
	// By default, sends and receives block until the other side is ready.
	// This allows goroutines to synchronize without explicit locks or condition variables.

	fmt.Println(x, y, x+y)
}
