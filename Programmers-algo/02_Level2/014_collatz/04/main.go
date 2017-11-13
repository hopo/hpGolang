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

func collats(num int) int {
	data := gen(num)
	if data == nil {
		return -1
	}
	return len(data) - 1
}

func gen(n int) []int {
	d := []int{n}
	for {
		l := len(d)
		if l > 500 {
			return nil
		}
		t := d[l-1]
		if t == 1 {
			return d
		}
		x := t % 2
		switch x {
		case 0:
			d = append(d, t/2)
		case 1:
			d = append(d, t*3+1)
		}
	}
	return d
}
