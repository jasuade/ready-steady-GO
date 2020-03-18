package main

import (
	"fmt"
)

//Vertex is a type struct, Go does not have clases, but you can define Methods on types.
type Vertex struct {
	x, y float64
}

//Abs . . .Method is a function with a special 'receiver' argument, in this case Vertex
//Here Abs method has a receiver of type Vertex named v
func (v Vertex) Abs() float64 {
	return v.x + v.y
}

//But remember that method is just a function with a receiver arg
func Abs2(v Vertex) float64 {
	return v.x + v.y
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(Abs2(v))

}
