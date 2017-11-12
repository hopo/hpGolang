package main

import (
	"fmt"
)

func main() {
	ex1 := getMinMax([]int{1, 2}, []int{3, 4}) // 10
	ex2 := getMinMax([]int{4, 2}, []int{3, 9}) // 30
	ex3 := getMinMax([]int{2, 3}, []int{5, 7}) // 29
	fmt.Println(ex1)
	fmt.Println(ex2)
	fmt.Println(ex3)

}

func getMinMax(a, b []int) int {
	var data []int
	for _, v := range a {
		for _, w := range b {
			data = append(data, v*w)
		}
	}

	l := len(data)
	min := data[0] + data[l-1]
	for i := 0; i < l/2; i++ {
		if min > data[i]+data[l-1-i] {
			min = data[i] + data[l-1-i]
		}
	}
	return min
}
