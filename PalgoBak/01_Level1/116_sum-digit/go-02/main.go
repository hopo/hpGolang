package main

import (
	"fmt"
)

func main() {
	ex1 := sum_digit(123)       // 6
	ex2 := sum_digit(999999999) // 81
	fmt.Println(ex1)
	fmt.Println(ex2)

}

func sum_digit(n int) (ret int) {
	isl := int_convert_each(n)
	for _, v := range isl {
		ret += v
	}
	return
}

func int_convert_each(n int) []int {
	di := 1
	for {
		if n-di < 0 {
			di = di / 10
			break
		}
		di *= 10
	}
	var isl []int
	for {
		if di == 1 {
			isl = append(isl, n)
			break
		}
		z := n / di
		isl = append(isl, z)
		n = n % di
		di = di / 10
	}
	return isl
}
