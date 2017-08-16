package main

import (
	"fmt"
)

func indexPower(is []int, n int) int {
	if n < 0 {
		return n
	}
	res := 1
	for i := 0; i < n; i++ {
		res *= is[n]
	}
	return res
}

func main() {
	fmt.Println(indexPower([]int{1, 2, 3, 4}, 2))    //9
	fmt.Println(indexPower([]int{1, 3, 10, 100}, 3)) //1000000
	fmt.Println(indexPower([]int{0, 1}, 0))          //1
	fmt.Println(indexPower([]int{1, 2}, -1))         //-1
}
