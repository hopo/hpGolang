package main

import (
	"fmt"
)

func main() {
	ex1 := caesar("a B z", 4)   // "e F d"
	ex2 := caesar("A b Z g", 5) // "F g E l"
	fmt.Println(ex1)
	fmt.Println(ex2)
}

func caesar(s string, n int) string {
	bsl := []byte(s)
	for i, v := range bsl {
		if v != 32 {
			if v == 90 {
				v = 64
			}
			if v == 122 {
				v = 96
			}
			bsl[i] = v + byte(n)
		}
	}
	return string(bsl)
}
