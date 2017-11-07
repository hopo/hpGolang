package main

import (
	"fmt"
	"github.com/hopo/hpkg"
)

func main() {
	ex1 := rm_small([]int{4, 3, 2, 1}) // [4 3 2]
	ex2 := rm_small([]int{10, 8, 22})  // [10 8 22]
	fmt.Println(ex1)
	fmt.Println(ex2)
}

func rm_small(isl []int) (ret []int) {
	var sm int
	for i := 1; i < len(isl); i++ {
		if isl[sm] > isl[i] {
			sm = i
		}
	}
	ret = hpkg.DelIsl(isl, sm)
	return
}
