package main

import (
	"fmt"
)

func main() {
	ex1 := foo01(3, 5)
	ex2 := foo02(3, 5)
	fmt.Println(ex1, ex2)
}

func foo01(a, b int) int {
	return func(a, b int) int {
		return a * b
	}(a, b)
}

func foo02(a, b int) int {
	r := func(a, b int) int {
		return a * b
	}
	return r(a, b)
}
