package main

import "fmt"

//A Type Assertion provides access to an interface values
//SO is used to check if the interface actually is implemented with the value
func main() {
	var i interface{} = "Hello"

	s := i.(string) //indicates if the interface has a value string and which one
	fmt.Println(s)

	s, ok := i.(string) //Indicates the value and true flag
	fmt.Println(s, ok)

	f, ok := i.(float64) // A the interface does not have any float value ok= false f=0
	fmt.Println(f, ok)

	f = i.(float64) //Throws an error because the interface does not have a float value
	fmt.Println(f)
}
