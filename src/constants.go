package main

import "fmt"

const Pi = 3.14 //could be characters, strings, bools or nums and can not be declared with :=

func main() {
	const World = "世界" //can be declared inside or outside a function but, again, not with :=
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

}
