//Exercise --> Given any number if:

// has 3 as a factor, add 'Pling' to the result.
// has 5 as a factor, add 'Plang' to the result.
// has 7 as a factor, add 'Plong' to the result.
// does not have any of 3, 5, or 7 as a factor, the result should be the digits of the number.

// Examples:

// 28 has 7 as a factor, but not 3 or 5, so the result would be "Plong".
// 30 has both 3 and 5 as factors, but not 7, so the result would be "PlingPlang".
// 34 is not factored by 3, 5, or 7, so the result would be "34".

package raindrops

import (
	"fmt"
	"strconv"
)

func raindrops(num int) string {
	fmt.Println("This is a raindrops exercise")
	resul := ""
	if num%3 == 0 {
		resul = "Pling"
	}
	if num%5 == 0 {
		resul = resul + "Plang"
	}
	if num%7 == 0 {
		resul = resul + "Plong"
	}
	if resul == "" {
		resul = strconv.Itoa(num)
	}
	return resul

}
