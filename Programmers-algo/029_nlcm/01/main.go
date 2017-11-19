package main

import (
	"fmt"
)

func main() {
	ex1 := nlcm([]int{7, 8, 2})
	ex2 := nlcm([]int{2, 6, 8, 14}) // 168
	fmt.Println(ex1)
	fmt.Println(ex2)
}

func nlcm(nums []int) int {
	var isl []int
	for _, v := range nums {
		box := []int{}
		box = denominator(v)
		isl = append(isl, box...)
	}
	for i, _ := range isl {
		for j, v := range isl {
			if i != j && isl[i] == v {
				isl[i] = 1
			}
		}
	}
	ret := 1
	for _, v := range isl {
		ret *= v
	}
	if ret == 1 {
		for _, v := range nums {
			ret *= v
		}
	}
	return ret
}

func denominator(n int) []int {
	var isl []int
	for i := 2; i < n; i++ {
		if n%i == 0 {
			isl = append(isl, i)
		}
	}
	if isl == nil {
		isl = append(isl, n)
	}
	return isl
}
