package main

import "fmt"

func main() {
	i, j := 42, 2701
	v := 56

	p := &i         //Pointer to variable i (42)
	fmt.Println(*p) //Show the content of the pointer p --> i = 42 know as "dereferencing" or "indirecting"
	*p = v          //When using * you set the VALUE of i throug the pointer
	fmt.Println(i)
	fmt.Println(*p)

	p = &j       //Here you change the direction of the pointer, now p --> j = 2701
	*p = *p / 37 //Again, using * you change the content of the pointer p--> j = 73
	fmt.Println(j)
}
