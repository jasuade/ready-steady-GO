package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var i Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{2, 3}

	i = f //now f (MyFloat) implements a(Abser)
	fmt.Println(i.Abs())

	i = &v // now *Vertex implements Abser
	fmt.Println(i.Abs())

	//i = v // Here v is a Vertex (new) but nor a *Vertex, thus DOES NOT implement Abser
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
