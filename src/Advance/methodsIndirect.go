package main

import "fmt"

type Vertex struct {
	x, y float64
}

func (v Vertex) Sum() float64 {
	return v.x + v.y
}

func SumFunction(v Vertex) float64 {
	return v.x + v.y
}

func main() {
	v := Vertex{2, 3}
	fmt.Println(v.Sum()) //  OK
	fmt.Println(SumFunction(v))
	//fmt.Println(Sum(v))  // OK
	//fmt.Println(Sum(&v)) // Compile error!

	p := &Vertex{4, 3}
	fmt.Println(p.Sum()) // OK
	//n this case, the method call p.Sum() is interpreted as (*p).Sum()
	fmt.Println(SumFunction(*p))
}
