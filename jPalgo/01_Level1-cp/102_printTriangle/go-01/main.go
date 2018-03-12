package main

import (
	"fmt"
	"github.com/hopo/hpkg"
)

func main() {
	ex1 := printTriangle(3)
	fmt.Println(ex1)
}

func printTriangle(num int) (rslt string) {
	for i := 1; i < num; i++ {
		rslt += hpkg.Smulti("*", i) + "\n"
	}
	rslt += hpkg.Smulti("*", num)
	return
}
