package main

import (
	"fmt"
)

func main() {
	ex1 := easy_unpack([]int{1, 2, 3, 4, 5, 6, 7, 9}) // [1 3 7]
	fmt.Println(ex1)
	ex2 := easy_unpack([]int{1, 1, 1, 1}) // [1 1 1]
	fmt.Println(ex2)
	ex3 := easy_unpack([]int{6, 3, 7}) // [6 7 3]
	fmt.Println(ex3)
}

func easy_unpack(nums []int) []int {
	return []int{nums[0], nums[2], nums[len(nums)-2]}
}
