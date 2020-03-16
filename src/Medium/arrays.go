package main

import "fmt"

//var num [10]int  //To init an empty array
var primes = [5]int{1, 3, 7, 11, 17} //Para inicializar el array tienes que poner el =

func main() {
	num := [12]float32{3.4, 6.7, 4.8, 0.45}
	var array [2]string //Arrays cannot be resized
	array[0] = "Holi"
	array[1] = "Mundi"

	fmt.Println(array)
	fmt.Println(array[1], array[0])
	fmt.Println(primes)
}
