package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	///you can not range over slices
	sl := make([][]uint8, dy) //create an empty slice of slices
	for y := 0; y < dy; y++ { //iterate over it
		sl[y] = make([]uint8, dx) //create the slice within the slices
		for x := 0; x < dx; x++ {
			sl[y][x] = uint8((x * y) / 2) //fulfill with values depending on the function
		}
	}
	return sl
}

func main() {
	pic.Show(Pic)
}
