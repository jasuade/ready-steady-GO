package main

import "fmt"

func main() {
	var i interface{} // The if is empty thus print (value, type) = (<nil>, <nil>)
	describe(i)

	i = 42
	describe(i) // The if is (42,42)

	i = "Hallo"
	describe(i) // The if is (Hallo, Hallo)

}

func describe(i interface{}) {
	fmt.Println("(%v, %T)\n", i, i)
}
