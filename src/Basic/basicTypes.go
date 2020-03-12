package main

import (
	"fmt"
	"math/cmplx"
)

var ( //variable declaration can be factored into blocks
	ToBe   bool       = false
	MaxInt int        = 1<<64 - 1 //int, uint and uintptr are usually 32/64 bits depending if the system is 32/64bits
	z      complex128 = cmplx.Sqrt(-5 + 12i)
) //When you need an integer value you should use int unless you have a specific reason to use a sized or unsigned integer type.

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe) //note the `%T` to print the value
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
