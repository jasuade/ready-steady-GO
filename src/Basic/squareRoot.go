package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	//For a given number x return the number z so that z^2 =
	resul := float64(0)
	// z := float64(1)
	// for z = (z*z - x) / (2*z); z <= x; z += 0.01
	for z := 1.0; z < x; z += 0.1 {
		if y := z * z; y < x {
			resul = z
		}
	}
	return resul
}

func main() {
	fmt.Println(Sqrt(7))
	fmt.Println(Sqrt(7) * Sqrt(7))
}
