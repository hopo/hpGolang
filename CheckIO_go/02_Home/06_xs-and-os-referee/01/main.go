package main

import (
	"fmt"
)

func xoReferee(data []string) string {
	var dt []string
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			dt = append(dt, string(data[i][j]))
		}
	}
	var os, xs, dot int
	for _, d := range dt {
		if "O" == d {
			os++
		}
		if "X" == d {
			xs++
		}
		if "." == d {
			dot++
		}
	}
	if dot == 0 {
		return "D"
	}
	if os > xs {
		return "O"
	}
	if os < xs {
		return "X"
	}
	return dt[4]
}

func main() {
	fmt.Println(xoReferee([]string{"X.O",
		"XX.",
		"XOO"})) //"X", "Xs wins"
	fmt.Println(xoReferee([]string{"OO.",
		"XOX",
		"XOX"})) //"O", "Os wins"
	fmt.Println(xoReferee([]string{"OOX",
		"XXO",
		"OXX"})) //"D", "Draw"
	fmt.Println(xoReferee([]string{"O.X",
		"XX.",
		"XOO"})) //"X", "Xs wins again"
}
