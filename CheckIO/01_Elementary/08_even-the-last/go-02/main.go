package main

import "fmt"

func evenLast(nums []int) int {
	var res int
	if len(nums) != 0 {
		for i := 0; i < len(nums); i += 2 {
			res += nums[i]
		}
		return res * nums[len(nums)-1]
	}
	return res
}

func main() {
	fmt.Println(evenLast([]int{0, 1, 2, 3, 4, 5})) //30, "(0+2+4)*5=30"
	fmt.Println(evenLast([]int{1, 3, 5}))          //30, "(1+5)*5=30"
	fmt.Println(evenLast([]int{6}))                //36, "(6)*6=36"
	fmt.Println(evenLast([]int{}))                 //0, "An empty array = 0"
}
