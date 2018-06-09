package main

import (
	"fmt"
)

func main() {
	var x, y, z int
	x, y, z = 2, 4, 8

	// integer data
	fmt.Println("x, y, z: ", x, y, z)

	// bitwise operator (OR)
	fmt.Printf("%v | %v(%04b OR %04b): ", x, z, x, z)
	fmt.Print(x | z)
	fmt.Printf("(%04b)\n", x|z)

	// bitwise operator (XOR)
	fmt.Printf("%v ^ %v(%04b XOR %04b): ", z, y, z, y)
	fmt.Print(z ^ y)
	fmt.Printf("(%04b)\n", z^y)

	// 8 - 2 = 6
	// 8 + 8 = 16 - 10 = 6
}
