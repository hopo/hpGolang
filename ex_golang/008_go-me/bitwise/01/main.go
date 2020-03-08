package main

import (
	"fmt"
)

func main() {
	data := []int{2, 3, 5, 6, 7, 8, 10}

	fmt.Print("data: ")
	for _, v := range data {
		fmt.Print(v, ", ")
	}

	fmt.Print("\n", "odd: ")
	for _, v := range data {
		if v%2 == 1 {
			fmt.Print(v, ", ")
		}
	}

	fmt.Print("\n", "even: ")
	for _, v := range data {
		if (v & 1) == 0 { // bitwise operate
			fmt.Print(v, ", ")
		}
	}
}
