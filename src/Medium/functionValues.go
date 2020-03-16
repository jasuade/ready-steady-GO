package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))      // Calls the function and execures sqrt(5*5 + 12*12)
	fmt.Println(compute(hypot))    //Calls the function compute returns 3, 4 and calls hypot to return sqrt(3*3 + 4*4)
	fmt.Println(compute(math.Pow)) //Calls the function math.Pow with the result of compute 3^4

}
