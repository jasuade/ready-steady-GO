package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {
	mapa := make(map[string]Vertex)
	mapa["Bell Labs"] = Vertex{40.45454, -74.264} // mapa --> {"Bell Labs":{40.45454, -74.264}}
	fmt.Println(mapa["Bell Labs"])

	//Map literals are like structs literals, but keys are required

	m := map[string]Vertex{
		"Bell Labs": Vertex{ //You can omit the top-leve type
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}

	fmt.Println(m)
}
