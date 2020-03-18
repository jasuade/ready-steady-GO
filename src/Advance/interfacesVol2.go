package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

//This method means that T struct implments the interface I,
//but is not literally specified
// func (t T) M() {
// 	fmt.Println(t.S)
// }

//In this case is the same than above but with a pointer receiver
func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var interfaz I = &T{"hallo"} //Creates the interface and assign the implementation with pointer
	//var interfaz I = T{"hallo"} Without pointer receiver
	describe(interfaz)
	interfaz.M()

	interfaz = F(math.Pi)
	describe(interfaz)
	interfaz.M()
}

func describe(i I) { //Depict the value of the interface and the type
	fmt.Printf("(%v, %T)\n", i, i)
}
