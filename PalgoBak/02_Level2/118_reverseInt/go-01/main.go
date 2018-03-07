package main

import (
	"fmt"
)

func main() {
	ex1 := reverseInt(118372) // 873211
	fmt.Println(ex1)
}

func reverseInt(n int) int {
	isl := split_int_each(n)
	sorted := sortRv_isl(isl)
	ret := each_int_makeNum(sorted)

	return ret
}

func split_int_each(n int) []int {
	numlen := 10
	for {
		if n-numlen < 0 {
			numlen /= 10
			break
		}
		numlen *= 10
	}

	var ret []int
	for {
		x := n / numlen
		ret = append(ret, x)
		r := n % numlen
		n = r
		numlen /= 10
		if numlen < 1 {
			break
		}
	}
	return ret
}

func sortRv_isl(isl []int) []int {
	for i, _ := range isl {
		for j, _ := range isl {
			if i != j && isl[i] > isl[j] {
				isl[i], isl[j] = isl[j], isl[i]
			}
		}
	}
	return isl
}

func each_int_makeNum(isl []int) int {
	l := len(isl)
	numlen := pow(10, l-1)
	var ret int
	for _, v := range isl {
		ret += v * numlen
		numlen /= 10
	}
	return ret
}

func pow(n, p int) int {
	ret := 1
	for i := 0; i < p; i++ {
		ret *= n
	}
	return ret
}
