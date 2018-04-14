package main

import (
	"fmt"
)

func radiationSearch(data [][]int) (ret []int) {
	return
}

func main() {

	d1 := [][]int{
		{1, 2, 3, 4, 5},
		{1, 1, 1, 2, 3},
		{1, 1, 1, 2, 2},
		{1, 2, 2, 2, 1},
		{1, 1, 1, 1, 1}} // [14, 1], "14 of 1"
	d2 := [][]int{
		{2, 1, 2, 2, 2, 4},
		{2, 5, 2, 2, 2, 2},
		{2, 5, 4, 2, 2, 2},
		{2, 5, 2, 2, 4, 2},
		{2, 4, 2, 2, 2, 2},
		{2, 2, 4, 4, 2, 2}} // [19, 2], "19 of 2"
	fmt.Pritnln(radiationSearch(d1))
	fmt.Pritnln(radiationSearch(d2))
}
