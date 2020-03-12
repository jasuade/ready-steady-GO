package main

import "fmt"

func main() {
	sum := 0
	//no parenthesis but braces are mandatory...
	for i := 0; i < 10; i++ { //the variables declared here are only visible within the for scope
		sum += i
	}
	fmt.Println(sum)

	//The init and incremental condition are optional, so you can
	for sum < 1200 {
		sum += sum
	}
	fmt.Println(sum)

	//Actually you can drop semicolons and get a while loop
	//And to create an infinite loop ypu can
	//for {}

}
