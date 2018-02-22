package main

import (
	"fmt"
	"github.com/hopo/hpkg"
)

func main() {
	ex1 := number_base("AF", 16) // 175
	fmt.Println(ex1)
	ex2 := number_base("101", 8) // 65
	fmt.Println(ex2)
}

func number_base(s string, n int) int {
	return hpkg.Number_base(s, n)
}
