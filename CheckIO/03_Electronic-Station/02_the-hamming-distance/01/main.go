package main

import (
	"fmt"
)

func main() {
	fmt.Println(hammingDistance(117, 17)) //3, "First example"
	fmt.Println(hammingDistance(1, 2))    //2, "Second example"
	fmt.Println(hammingDistance(16, 15))  //5, "Third example"
	fmt.Println(hammingDistance(256, 8))  //2, "hp ex""
}

func hammingDistance(n, m int) (ret int) {
	x := bConverter(n)
	y := bConverter(m)

	for i := 0; i < len(x); i++ {
		if x[i] != y[i] {
			ret++
		}
	}

	return
}

func bConverter(n int) string {
	s := fmt.Sprintf("%09b", n)
	return s
}
