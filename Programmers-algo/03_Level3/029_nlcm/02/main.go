package main

import (
	"fmt"
)

func main() {
	ex1 := nlcm([]int{3, 5, 6, 7, 8, 9, 15}) // 2520
	ex2 := nlcm([]int{2, 6, 8, 14})          // 168
	fmt.Println(ex1)
	fmt.Println(ex2)
}

func nlcm(nums []int) int {
	// *** find max value in nums
	var max int
	max = nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}

	// *** power max: max*2 max*3... if remainder is 0 by all values divide that is lcm
	var smpl, count int
	for i := 1; ; i++ {
		count = 0
		smpl = max * i
		for _, v := range nums {
			if smpl%v == 0 {
				count++
			}
		}
		if count == len(nums) {
			return smpl
		}
	}
	return -1
}
