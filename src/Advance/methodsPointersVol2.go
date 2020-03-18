package main

import "fmt"

type Vertex struct {
	x, y float64
}

//There are two main reasons to use pointers instead of values in the methods
// 1. The method can modify directly the value that its receiver pointers to
// 2. To avoid copying the value on each method call. This is even more efficient
//if the struct is large

func (v *Vertex) Scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func (v *Vertex) Sum() float64 {
	return v.x + v.y
}

func main() {
	v := &Vertex{2, 3}
	fmt.Printf("Before scaling: %+v, Sum: %v\n", v, v.Sum()) //when printing structs, the plus flag (%+v) adds field names
	v.Scale(10)
	fmt.Printf("Afer scaling: %+v, Sum: %v\n", v, v.Sum())
	fmt.Println("Without using the sum: ", v)
}

// In general, all methods on a given type should have
// ither value or pointer receivers, but not a mixture of both. (We'll see why over the next few pages.)
