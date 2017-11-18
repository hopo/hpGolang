package main

import (
	"fmt"
)

func main() {
	ex1 := nextBigNumber(78) // 83
	ex2 := nextBigNumber(47) // 55
	fmt.Println(ex1)
	fmt.Println(ex2)
}

func nextBigNumber(n int) int {
	ns := binary_convert(n)
	one := ea_one(ns)
	for i := n + 1; ; i++ {
		s := binary_convert(i)
		o := ea_one(s)
		if one == o {
			return i
		}
	}
	return -1
}

func ea_one(s string) int {
	var ret int
	bsl := []byte(s)
	for _, v := range bsl {
		if v == 49 {
			ret++
		}
	}
	return ret
}

func binary_convert(n int) string {
	//fmt.Println("n int:", n) // checker
	//fmt.Printf("%b\n", n) // checker
	var tsl []int
	for i := 1; i < n+1; i *= 2 {
		tsl = append([]int{i}, tsl...) // push front
	}

	zno := make([]int, len(tsl))
	for i := 0; i < len(tsl); i++ {
		n = n - tsl[i]
		if n == 0 {
			zno[i] = 1
			break
		}
		if n > 0 {
			zno[i] = 1
		}
		if n < 0 {
			n = n + tsl[i]
		}
	}

	var r string
	for _, v := range zno {
		r += string(v + 48) // int make string
	}
	return r
}
