package main

import (
	"fmt"
)

func main() {
	ex1 := collats(6)  // 8
	ex2 := collats(11) // 14
	fmt.Println(ex1)
	fmt.Println(ex2)
}

// 6→3→10→5→16→8→4→2→1

var data []int

func collats(num int) int {
	data = []int{}
	gen(num)
	return len(data) - 1
}

func gen(num int) int {
	data = append(data, num)
	if len(data) > 500 {
		return -1
	}
	if num == 1 {
		return 1
	}
	x := num % 2
	switch x {
	case 0:
		num = num / 2
	case 1:
		num = num*3 + 1
	}
	return gen(num)
}
