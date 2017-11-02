package main

import (
	"fmt"
	"math"
)

func main() {
	ex1 := nextSqaure(121) // 144, 12*12
	ex2 := nextSqaure(4)   // 9, 3*3
	ex3 := nextSqaure(5)   // 0, "no"
	fmt.Println(ex1)
	fmt.Println(ex2)
	fmt.Println(ex3)
}

func nextSqaure(n int) (rslt int) {
	x := math.Sqrt(float64(n))
	f := math.Pow(x, 2)
	if f == float64(n) {
		rslt = int(math.Pow(x+1, 2))
	} else {
		rslt = -1
	}
	return
}
