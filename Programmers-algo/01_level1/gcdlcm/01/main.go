package main

import (
	"fmt"
	"spila"
)

func main() {
	ex1 := spila.Gcdlcm(3, 12) // [3 12]
	ex2 := spila.Gcdlcm(4, 7)  // [1, 28]
	fmt.Println(ex1)
	fmt.Println(ex2)
}
