package main

import (
	"fmt"
	"math"
)

func absoluteSorting(is []int) []int {
	l := len(is)
	for i := 1; i < l; i++ {
		for j := 0; j < l-1; j++ {
			if math.Abs(float64(is[i])) < math.Abs(float64(is[j])) {
				is[i], is[j] = is[j], is[i]
			}
		}
	}
	return is
}

func main() {
	fmt.Println(absoluteSorting([]int{-20, -5, 10, 15})) //[-5, 10, 15, -20]
	fmt.Println(absoluteSorting([]int{1, 2, 3, 0}))      //[0, 1, 2, 3]
	fmt.Println(absoluteSorting([]int{-1, -2, -3, 0}))   //[0, -1, -2, -3]
}
