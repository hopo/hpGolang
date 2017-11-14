package main

import (
	"fmt"
)

func main() {
	ex1 := evenOrOdd(3) // Odd
	ex2 := evenOrOdd(2) // Even

	fmt.Println(ex1)
	fmt.Println(ex2)
}

func evenOrOdd(num int) (rslt string) {
	if num%2 == 0 {
		rslt = "Even"
	} else {
		rslt = "Odd"
	}
	return
}
