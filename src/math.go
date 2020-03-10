//Playing with math and imports

package main

//the imports could be included in braces and like that:
//import "fmt"
//import "math"

import (
	"fmt"
	"math"
)

func main() {
	//NOTE the PrintF for the display of variables, instead of Println
	fmt.Printf("The square root of 7 is %g, I think.\n", math.Sqrt(7))
}
