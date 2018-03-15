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
	isl := []int{2, 3}
	for i := 2; i < n+1; i++ {
		if i%2 != 0 && i%3 != 0 {
			isl = append(isl, i)
		}
	}
	return len(isl)
}
