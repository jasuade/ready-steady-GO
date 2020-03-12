package main

import "fmt"

const (
	//Create a huge number shifting 1 bit left 100 places to binary num created by 1 followed by 100 zeroes
	Big = 1 << 100
	//Shift 99 positions to the right to ger 1 << 1 = 10, or 2
	Small = Big >> 99
)

func needInt(x int) int           { return x*10 + 1 } //func that receives an int and return that int multiplied by 10 and + 1
func needFloat(x float64) float64 { return x * 0.1 }  //func that receives a float and returns that num multiplied by 0,1

func main() {
	fmt.Println(needInt(Small))
	//fmt.Println(needInt(Big)) //this creates an int to big that overflows
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
