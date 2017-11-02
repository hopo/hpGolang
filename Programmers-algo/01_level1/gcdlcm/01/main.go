package main

import (
	"fmt"
)

func main() {
	ex1 := gcdlcm(3, 12) // [3 12]
	fmt.Println(ex1)
}

func gcdlcm(a, b int) (rslt []int) {
	// gcd, lcm
	fa := facto(a)
	fb := facto(b)
	rslt = append(rslt, a, b)
	return
}

func facto(a int) (rslt []int) {
	go func(a int) int {
		if a%2 == 0 {
			rslt = append(rslt, 2)
			a = a / 2
			return
		}
	}(a)
	return
}
