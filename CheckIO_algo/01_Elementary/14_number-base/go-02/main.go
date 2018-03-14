package main

import (
	"fmt"
	"math"
)

func main() {
	ex1 := number_base("AF", 16) // 175
	fmt.Println(ex1)
	ex2 := number_base("101", 2) // 5
	fmt.Println(ex2)
}

func number_base(s string, n int) int {
	bsl := []byte(s)
	var isl []int
	for _, v := range bsl {
		if 64 < v && v < 91 {
			isl = append(isl, int(v-55))
		}
		if 47 < v && v < 59 {
			isl = append(isl, int(v-48))
		}
	}

	var j int
	var ret int
	for i, _ := range isl {
		j = len(isl) - 1 - i
		mp := math.Pow(float64(n), float64(i))
		ret += isl[j] * int(mp)
	}
	return ret
}
