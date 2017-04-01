package main

import (
	"fmt"
	"sort"
)

func main() {
	data := []int{1, 2, 4, 9, 9, 10}
	sort.Ints(data)

	for x, y := range data {
		//i := sort.Search(len(data), func(i int) bool { return data[i] >= x })
		//j := sort.Search(len(data), func(j int) bool { return data[j] > x })
		//fmt.Println(x, y, ":", i, j)
	}
	//fmt.Println(data, i, j)
}
