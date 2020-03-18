package main

import "fmt"

func do(i interface{}) { //Practical uses of an interface! a method can accept many data types
	switch v := i.(type) { //Checkt the type underneath the interface
	case int: //v has a type int
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string: //v has a type string
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
