package main

import (
	"fmt"
)

func main() {
	board1 := [][]int{{1, 2, 3, 5}, {5, 6, 7, 8}, {4, 3, 2, 1}} // 16, (5)+(7)+(4)
	ex1 := hopscotch(board1, 3)
	fmt.Println(ex1)
	board2 := [][]int{{1, 2, 7, 5}, {3, 2, 6, 4}, {2, 3, 4, 5}, {1, 1, 9, 1}} // 15, (7)+(4)+(4)+(1)
	ex2 := hopscotch(board2, 4)
	fmt.Println(ex2)
}

type Data struct {
	sum   int
	avoid int
}

func hopscotch(board [][]int, size int) int {
	dt := Data{0, -1} // init
	for i := 0; i < size; i++ {
		dt = sum(dt, board[i])
	}
	return dt.sum
}

func sum(dt Data, isl []int) Data {
	var s, a int
	for i, v := range isl {
		if dt.avoid != i && s < v {
			s = v
			a = i
		}
	}
	var ret Data
	ret = Data{dt.sum + s, a}
	return ret
}
