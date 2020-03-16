package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("the value is:", m["Answer"])

	m["Answer"] = 73
	fmt.Println("the value is:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("the value is:", m["Answer"])

	v, ok := m["Answer"] //f key is in m, ok is true. If not, ok is false.
	//If key is not in the map, then elem is the zero value for the map's element type.
	fmt.Println("The Value:", v, "Present?", ok)
}
