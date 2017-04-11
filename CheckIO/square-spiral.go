package main

import (
	"fmt"
)

func findDistance(first, second int) int {
	fmt.Println(first, second)
	return 0
}

func main() {
	fmt.Println(findDistance(1, 9))   //2, "1st example"
	fmt.Println(findDistance(9, 1))   //2, "2nd example"
	fmt.Println(findDistance(10, 25)) //1, "3rd example"
	fmt.Println(findDistance(5, 9))   //4, "4th example"
	fmt.Println(findDistance(26, 31)) //5, "5th example"
	fmt.Println(findDistance(50, 16)) //10, "6th example"
}
