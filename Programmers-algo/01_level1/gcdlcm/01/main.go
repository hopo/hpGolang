package main

import (
	"fmt"
	"github.com/hopo/hpkg"
)

func main() {
	ex1 := hpkg.Gcdlcm(3, 12) // [3 12]
	ex2 := hpkg.Gcdlcm(4, 7)  // [1, 28]
	fmt.Println(ex1)
	fmt.Println(ex2)
}
