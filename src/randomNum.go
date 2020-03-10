//Function that generates X random number, being seed by the clock
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		fmt.Println("Iteration", i)
		fmt.Println("Random number", rand.Intn(500))
	}

}
