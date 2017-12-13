package main

import (
	"fmt"
)

func main() {
	ex := sumMatrix([][]int{{1, 2}, {2, 3}}, [][]int{{3, 4}, {5, 6}})
	fmt.Println(ex) // [[4 6] [7 9]]
}

func sumMatrix(d1, d2 [][]int) (ret [][]int) {
	ret = [][]int{{0, 0}, {0, 0}}
	for i := 0; i < len(d1); i++ {
		for j := 0; j < len(d1[0]); j++ {
			ret[i][j] = d1[i][j] + d2[i][j]
		}
	}
	return
}
