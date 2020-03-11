package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x)) //Sprint formats the string and does not print to console. It returns the formatted string.
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v //the variable v is declared in the if so only valid in this scope, also in the else
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	fmt.Println(sqrt(2), sqrt(-1))
	fmt.Println(pow(3, 2, 10), //pow 3^2 < 10 so it's return
		pow(3, 3, 20)) //3^3
}
