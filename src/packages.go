//Random number generator
package main

import (
	"fmt"
	"math/rand"
	"time"
)

//This function selects a random number between 0 and x,
//But to create a totally random number each execurion you have to use Seed
//and feed the entrance with a random value.
func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Choose a random number", rand.Intn(500))
}
