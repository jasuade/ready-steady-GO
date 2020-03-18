package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	x, y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.x + v.y)
}

func (v *Vertex) Scale(f float64) { //With the pointer receiver you modify the original Vertex
	v.x = v.x * f
	v.y = v.y * f
	// With a value receiver (no *), the Scale method operates on a copy of the original Vertex value.
}

func main() {
	set := Vertex{3, 2}
	set.Scale(10)
	fmt.Println(set)
	fmt.Println(set.Abs())
}
