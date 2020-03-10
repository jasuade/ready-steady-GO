//In Go a Name or method is exported when it begins in Captal letter
//and thus, is internal when begins in lowecase

package main

import (
	"fmt"
	"math"
)

//Note that Pi is exported thus in captal letter
//using math.pi would not work
func main() {
	fmt.Println(math.Pi)
}
