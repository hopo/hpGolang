package main

import (
	"fmt"
)

func main() {
	ex := adder(3, 5) // 12, 3+4+5
	fmt.Println(ex)
}

func adder(a, b int) int {
	if a == b {
		return a
	}
	if a > b {
		a, b = b, a
	}
	var ret int
	for i := a; i < b+1; i++ {
		ret += i
	}
	return ret
}
