package main

import (
	"fmt"
)

func main() {
	ex1 := second_index("sims", "s") // 3, "First"
	fmt.Println(ex1)
	ex2 := second_index("find the river", "e") // 12, "Second"
	fmt.Println(ex2)
	ex3 := second_index("hi", " ") // -1, "Third"
	fmt.Println(ex3)
	ex4 := second_index("hi mayor", " ") // -1, "Fourth"
	fmt.Println(ex4)
	ex5 := second_index("hi mr Mayor", " ") // 5, "Fifth"
	fmt.Println(ex5)
}

func second_index(text string, symbol string) int {
	tsl := []byte(text)
	sym := []byte(symbol)
	var box []int
	for i, v := range tsl {
		if v == sym[0] {
			box = append(box, i)
		}
	}
	if len(box) < 2 {
		return -1
	}
	return box[1]
}
