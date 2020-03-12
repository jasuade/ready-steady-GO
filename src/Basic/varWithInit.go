package main

import "fmt"

var i, j int = 1, 2

//unable to do i,j := 1,2

func main() {
	k := 0                                    //:= is the short assigment to not use var but ONLY within functions, outside every statement begins with var func, etc.
	var c, python, java = true, false, "asdf" //if a value for initialize is presented you can omit the type
	fmt.Println(i, j, c, python, java)
}
