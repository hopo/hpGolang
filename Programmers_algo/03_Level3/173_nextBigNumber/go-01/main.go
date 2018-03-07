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
	ns := fmt.Sprintf("%b", n)
	one := ea_one(ns)
	for i := n + 1; ; i++ {
		s := fmt.Sprintf("%b", i)
		o := ea_one(s)
		if one == o {
			// fmt.Println(ns, "->", s) // for check
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
