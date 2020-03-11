package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	//For a given number x return the number z so that z^2 =
	resul := float64(0)
	for z := 1.0; z < x; z += 0.1 {
		if y := z * z; y < x {
			resul = z
		}
	}
	return resul
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(1.4 * 1.4)
}
