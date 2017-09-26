package main

import (
	"fmt"
	"strconv"
)

func main() {
	//fmt.Println(hammingDistance(117, 17)) //3, "First example"
	//fmt.Println(hammingDistance(1, 2))    //2, "Second example"
	//fmt.Println(hammingDistance(16, 15))  //5, "Third example"
	fmt.Println(hammingDistance(16, 15)) // "hp ex""
}

func hammingDistance(n, m int) int {
	x := itotypebConverter(n)
	y := itotypebConverter(m)

	fmt.Println("x, y, z:", x, y, z)

	ss := strconv.Itoa(z)
	var ii int
	for _, v := range ss {
		c, _ := strconv.Atoi(string(v))
		ii += c
	}

	var ret int = ii
	return ret
}

func itotypebConverter(n int) int {
	//s := fmt.Sprintf("%08b", n)
	s := strconv.FormatInt(int64(n), 2)
	fmt.Println("n:", s)

	i, _ := strconv.Atoi(s)
	return i
}

/*
func typebtoiConverter(n int) int {
	s := strconv.Itoa(n)

	var ii int
	for i := len(s) - 1; i > -1; i-- {
		ii += int(math.Pow(2, float64(i)))
	}
	var ret int = ii
	return ret
}
*/
