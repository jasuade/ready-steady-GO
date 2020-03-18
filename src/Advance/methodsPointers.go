package main

import "fmt"

type Vertex struct {
	x, y float64
}

func (v *Vertex) Scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func ScaleInFunction(v *Vertex, f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func main() {
	v := Vertex{3, 2}
	v.Scale(2)
	fmt.Println("Should be {6 and 4}", v)

	ScaleInFunction(&v, 10)
	fmt.Println("Should be {60 and 40}", v)

	p := &Vertex{4, 3}
	p.Scale(2)
	fmt.Println("Should be {8 and 6}", p)
	ScaleInFunction(p, 10)
	fmt.Println("Should be {80 and 60}", p)

}
