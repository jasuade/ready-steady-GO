package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	listOfFields := strings.Fields(s)
	fmt.Println(listOfFields)
	mapa := make(map[string]int)
	for x := 0; x < len(listOfFields); x++ {
		if mapa[listOfFields[x]] != 0 {
			mapa[listOfFields[x]] += 1
		} else {
			mapa[listOfFields[x]] = 1
		}
	}

	return mapa
}

func main() {
	wc.Test(WordCount)
}
