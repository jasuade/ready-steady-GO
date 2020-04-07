package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[:] {
		//fmt.Println(strings.Join(os.Args[], " "))
		fmt.Printf("Index %d and Value %s \n", i, arg)
	}

}
