// go 01/
package main

import (
	"fmt"
)

func hammingDistance(n, m int64) int {
	fmt.Printf("%08b\n", n)
	fmt.Printf("%08b\n", m)
	return 0
}

func main() {
	fmt.Println(hammingDistance(117, 17)) // "hp ex""
	//fmt.Println(hammingDistance(117, 17)) //3, "First example"
	//fmt.Println(hammingDistance(1, 2))    //2, "Second example"
	//fmt.Println(hammingDistance(16, 15))  //5, "Third example"
}
