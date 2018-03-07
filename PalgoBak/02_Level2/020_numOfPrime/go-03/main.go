package main

import (
	"fmt"
)

func main() {
	ex1 := numOfPrime(10) // 4, [2 3 5 7]
	ex2 := numOfPrime(5)  // 3, [2 3 5]
	fmt.Println(ex1)
	fmt.Println(ex2)
}

func numOfPrime(n int) int {
	r := makenop(n)
	return len(r)
}

func makenop(lmt int) []int {
	var isl []int
	for i := 2; i < lmt+1; i++ {
		isl = func(isl []int, i int) []int {
			for _, v := range isl {
				if i%v == 0 {
					return isl
				}
			}
			return append(isl, i)
		}(isl, i)
	}
	return isl
}
