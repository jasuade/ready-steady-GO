package main

import "fmt"

func main() {
	//Channels can be buffered so its more or less like a FIFO queue
	canal := make(chan int, 2)
	canal <- 1
	canal <- 2
	fmt.Println(<-canal)
	fmt.Println(<-canal)
}
