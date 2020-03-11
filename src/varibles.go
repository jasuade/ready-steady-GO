package main

import "fmt"

var c, java, python bool // var statement declares variables, a list of the same type in this case
//the var statement can be at package or function level and it inicilized at 0 or false or ''

func main() {
	var i int
	fmt.Println(i, c, python, java)
}
