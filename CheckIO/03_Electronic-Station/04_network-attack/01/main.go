package main

import (
	"fmt"
)

func capture(data [][]int) (ret int) {
	return
}

func main() {
	d1 := [][]int{
		{0, 1, 0, 1, 0, 1},
		{1, 8, 1, 0, 0, 0},
		{0, 1, 2, 0, 0, 1},
		{1, 0, 0, 1, 1, 0},
		{0, 0, 0, 1, 3, 1},
		{1, 0, 1, 0, 1, 2}} // 8, "Base example"
	d2 := [][]int{
		{0, 1, 0, 1, 0, 1},
		{1, 1, 1, 0, 0, 0},
		{0, 1, 2, 0, 0, 1},
		{1, 0, 0, 1, 1, 0},
		{0, 0, 0, 1, 3, 1},
		{1, 0, 1, 0, 1, 2}} // 4, "Low security"
	d3 := [][]int{
		{0, 1, 1},
		{1, 9, 1},
		{1, 1, 9}} // 9, "Small"

	fmt.Println(capture(d1))
	fmt.Println(capture(d2))
	fmt.Println(capture(d3))
}
