package main

import "fmt"

func split(sum int) (x, y int) 
{
	x = sum * 4 / 9
	y = sum - x
	return //naked return as the return values are defined in the functions's declaration
	// only to be used in short functions, to not harm readability
}

func main() {
	fmt.Println(split(17))
}
