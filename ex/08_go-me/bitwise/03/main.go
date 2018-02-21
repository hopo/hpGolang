package main

import (
	"fmt"
)

func main() {
	var x, y int
	x, y = 3, 4

	// 0011 -> 1100
	fmt.Printf("%v<<2: %v\n", x, x<<2)

	// 0011 -> 1100 + 0011 = 1111
	// 1100 + 0011 = 1111
	fmt.Printf("%v<<2+%v: %v\n", x, x, x<<2+x)

	// 0100 -> 0001
	fmt.Printf("%v>>1: %v\n", y, y>>1)
}
