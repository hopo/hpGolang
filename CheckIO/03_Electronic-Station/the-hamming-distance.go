package main

import (
	"fmt"
)

func hammingDistance(n, m int) int {
	fmt.Println(n, m)
	return 0
}

func main() {
	fmt.Println(hammingDistance(117, 17)) //3, "First example"
	fmt.Println(hammingDistance(1, 2))    //2, "Second example"
	fmt.Println(hammingDistance(16, 15))  //5, "Third example"
}
