// *** api lab

package main

import (
	"fmt"
)

func main() {
	ex1 := sort_isl([]int{1, 1, 8, 3, 7, 2})
	fmt.Println("sort_isl:", ex1)

	ex2 := isl_makeNum([]int{3, 6, 9})
	fmt.Println("isl_makeNum:", ex2)

	ex3 := pow(2, 3)
	fmt.Println("pow:", ex3)
}

// isl sort
func sort_isl(isl []int) []int {
	for i, _ := range isl {
		for j, _ := range isl {
			if i != j && isl[i] < isl[j] { // reverse sort. change '>'
				isl[i], isl[j] = isl[j], isl[i]
			}
		}
	}
	return isl
}

// each int in isl make Number
func isl_makeNum(isl []int) int {
	l := len(isl)
	numlen := pow(10, l-1) // math power function need
	var ret int
	for _, v := range isl {
		ret += v * numlen
		numlen /= 10
	}
	return ret
}

// math power n**p
func pow(n, p int) int {
	ret := 1
	for i := 0; i < p; i++ {
		ret *= n
	}
	return ret
}
