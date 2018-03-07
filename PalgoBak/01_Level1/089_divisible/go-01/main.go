package main

import (
	"fmt"
)

func main() {
	ex := divisible([]int{5, 9, 7, 10}, 5) // [5 10]
	fmt.Println(ex)
}

func divisible(isl []int, in int) (ret []int) {
	for _, v := range isl {
		if v%in == 0 {
			ret = append(ret, v)
		}
	}
	return
}
