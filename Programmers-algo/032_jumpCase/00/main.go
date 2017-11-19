package main

import (
	"fmt"
)

func main() {
	ex1 := jumpCase(5) // 5
	fmt.Println(ex1)
}

func jumpCase(n int) int {
	fmt.Println("n:", n)

	isl := big_dduim(n)
	fmt.Println("isl:", isl)

	return -1
}

func big_dduim(n int) []int {
	var isl []int
	for {
		if n == 2 {
			isl = append(isl, n)
			break
		}
		if n == 1 {
			isl = append(isl, n)
			break
		}
		if n-2 > 0 {
			isl = append(isl, 2)
			n -= 2
		}
	}
	return isl
}
