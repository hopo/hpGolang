package main

import (
	"fmt"
)

func main() {
	ex1 := digit_reverse(12345) // [5 4 3 2 1]
	ex2 := digit_reverse(321)   // [1 2 3]
	fmt.Println(ex1)
	fmt.Println(ex2)
}

func digit_reverse(n int) []int {
	isl := int_convert_each(n)
	var ret []int
	for i := 0; i < len(isl); i++ {
		j := len(isl) - 1 - i
		ret = append(ret, isl[j])
	}
	return ret

}

func int_convert_each(num int) []int {
	numlen := 10
	for {
		if num-numlen < 0 {
			numlen /= 10
			break
		}
		numlen *= 10
	}
	var ret []int
	var ea int
	for {
		ea = num / numlen
		ret = append(ret, ea)
		num %= numlen
		numlen /= 10
		if numlen < 1 {
			break
		}
	}
	return ret
}
