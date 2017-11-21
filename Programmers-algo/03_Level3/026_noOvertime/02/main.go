package main

import (
	"fmt"
)

func main() {
	//ex1 := noOvertime(4, []int{4, 3, 3}) // 12, (2**2)*3
	//ex2 := noOvertime(7, []int{9, 11})   // 85, (7*7)+(6*6)
	//fmt.Println(ex1)
	//fmt.Println(ex2)

	least_ot(4, []int{4, 3, 3})
}

func noOvertime(n int, works []int) int {
	// * (total works - n) / wparts ... remainder
	var total int
	wparts := len(works)
	for _, v := range works {
		total += v
	}
	ovtime := total - n
	div := ovtime / wparts
	rmnd := ovtime % wparts

	var ret, q int
	for i := 0; i < wparts; i++ {
		q = 0
		if rmnd != 0 {
			q = 1
			rmnd--
		}
		ret += (div + q) * (div + q)
	}

	return ret
}

func least_ot(n int, works []int) {
	var total int
	//wparts := len(works)
	for _, v := range works {
		total += v
	}
	//ov := total - n

	var ret int
	var x, y, z []int
	for i := 0; i < 7; i++ {
		x = []int{6, 5, 4, 3, 4, 3, 2}
		y = []int{0, 1, 2, 3, 1, 2, 2}
		z = []int{0, 0, 0, 0, 1, 1, 2}
		ret = (x[i] * x[i]) + (y[i] * y[i]) + (z[i] * z[i])
		fmt.Println(ret)
	}
}
