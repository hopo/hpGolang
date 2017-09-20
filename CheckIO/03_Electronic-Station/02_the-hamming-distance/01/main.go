package main

import (
	"fmt"
	"strconv"
)

func main() {
	//fmt.Println(hammingDistance(117, 17)) //3, "First example"
	//fmt.Println(hammingDistance(1, 2))    //2, "Second example"
	//fmt.Println(hammingDistance(16, 15))  //5, "Third example"
	fmt.Println(hammingDistance(117, 17)) // "hp ex""
}

func hammingDistance(n, m int) int {
	x := ntobConverter(n)
	y := ntobConverter(m)

	z := x - y

	var ret int
	ret = btonConverter(z)

	return ret
}

func ntobConverter(n int) int {
	//s := fmt.Sprintf("%08b", n)
	s := strconv.FormatInt(int64(n), 2)

	i, _ := strconv.Atoi(s)
	return i
}

func btonConverter(n int) int {
	s := strconv.FormatInt(int64(n), 10)
	i, _ := strconv.Atoi(s)
	fmt.Println("bton", i)

	return i
}
