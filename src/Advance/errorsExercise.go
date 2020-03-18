package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot Sqrt negative number: %f", e)
}

func Sqrt(x float64) (float64, error) {
	if x > 0 {
		resul := float64(0)
		z := float64(1)
		for z = (z*z - x) / (2 * z); z <= x; z += 0.01 {
			if y := z * z; y <= x {
				resul = z
			}
		}
		return resul, nil
	}
	err := ErrNegativeSqrt(x)
	return 0, err
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
