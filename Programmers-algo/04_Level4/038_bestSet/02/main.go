package main

import (
	"fmt"
)

func main() {
	//	ex1 := bestSet(3, 13) // [4 4 5]
	ex2 := bestSet(3, 11)
	//	fmt.Println(ex1)
	fmt.Println("Answer:", ex2)
}

func proof(n, s int) {
	box := make([]int, n)
	// hs := (s + 1) / 2 // s is int odd

	var x int
	x = 6
	var xx []int
	for i, _ := range box {
		xx = []int{s - (x * i), x / 2, x / 2}
		box[i] = xx[i]
		fmt.Println(box)
	}

	fmt.Println(box)
}

func bestSet(n, s int) []int {
	// proof 4*4*5 is most big number
	proof(n, s)
	var ret []int
	de := s / n
	r := s % n
	x := 1
	for i := 0; i < n; i++ {
		x = 0
		if r != 0 {
			x = 1
			r--
		}
		ret = append([]int{de + x}, ret...)
	}

	return ret
}
