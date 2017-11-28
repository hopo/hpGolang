package main

import (
	"fmt"
)

func main() {
	testBoard1 := [][]string{
		{"X", "O", "O", "O", "X"},
		{"X", "O", "O", "O", "O"},
		{"X", "X", "O", "O", "O"},
		{"X", "X", "O", "O", "O"},
		{"X", "X", "X", "X", "X"},
	} // 9
	ex1 := findLargestSquare(testBoard1)
	fmt.Println(ex1)

}

func findLargestSquare(board [][]string) interface{} {
	var ret [][]string
	var box []string
	for i, _ := range board {
		box = o_area(string(i+97), board[i])
		ret = append(ret, box)
	}
	return ret
}

func o_area(s string, row []string) []string {
	var box []string
	for i, v := range row {
		if v == "O" {
			box = append(box, s+string(48+i))
		}
	}
	return box
}

func str_plus_int(s string, n int) string {
	return s + string(n+48)
}
