package main

import (
	"fmt"
)

func main() {
	ex1 := number_generator(3, 5) // [3 6 9 12 15]
	ex2 := number_generator(4, 3) // [4 8 12]
	fmt.Println(ex1)
	fmt.Println(ex2)
}

func number_generator(x, n int) (ret []int) {
	for i := 1; i < n+1; i++ {
		ret = append(ret, x*i)
	}
	return
}
