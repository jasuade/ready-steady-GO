package main

import "fmt"

//This is used as Java classes
type Person struct {
	Name string
	Age  int
}

//Generic interface to print values
//A Stringer is a type that can describe itself as a string
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	//Instanciations of the classes
	a := Person{"Benito Kamelas", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)

}
