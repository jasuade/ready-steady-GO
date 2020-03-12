package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4 //when there is no specific type int, the variable is infered
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f) //to cast values except for int to string that we use strconv.Itoa(x)
	fmt.Println(x, y, z)
}
