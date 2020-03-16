package main

import "fmt"

type Estructura struct {
	x, y int
}

var (
	v1 = Estructura{1, 2}  // has type Estructura
	v2 = Estructura{x: 1}  // Y:0 is implicit
	v3 = Estructura{}      // X:0 and Y:0
	p  = &Estructura{1, 2} // has type *Estructura
)

func main() {

	fmt.Println(v1, v2, p, v3)
}
