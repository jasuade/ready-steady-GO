package main

import "fmt"

type vert struct { // structs with several fields
	x int
	y int
}

func main() {
	fmt.Println(vert{1, 2})
	//fmt.Println(vert.x) // you can not access directly to the filds of the struct because is like a data type
	v := vert{} // you have to assign a vale name for this struct
	v.x = 4     //Then you can access the methods and variable inside the stuct
	fmt.Println(v.x)

	//What about Pointers
	p := &vert{} //Both ways are valid to declare a pointer
	j := &v      //other way to do it
	p.x = 1e9    // the pointer access the struct and change the value
	fmt.Println(p)
	fmt.Println(j)

}
