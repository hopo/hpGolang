package main

import (
    "fmt"
)

func xoReferee(data []string) string {
	const i int = 0
	x, y, z := data[i], data[i+1], data[i+2]

	if x[i] == y[i+1] && y[i+1] == z[i+2] {
		return string(x[i])
	}
	if x[i+2] == y[i+1] && y[i+1] == z[i] {
		return string(x[i+2])
	}

	for j, D := range []string{x, y, z} {
		if D[i] == D[i+1] && D[i+1] == D[i+2] {
			return string(D[i])
		}
		if x[j] == y[j] && y[j] == z[j] {
			return string(x[j])
		}
	}

	return "D"
}

func main() {
    fmt.Println(xoReferee([]string{"X.O",
                                   "XX.",
                                   "XOO"})) //"X", "Xs wins"
    fmt.Println(xoReferee([]string{"OO.",
                                   "XOX",
                                   "XOX"}))	//"O", "Os wins"
    fmt.Println(xoReferee([]string{"OOX",
                                   "XXO",
                                   "OXX"})) //"D", "Draw"
    fmt.Println(xoReferee([]string{"O.X",
                                   "XX.",
                                   "XOO"})) //"X", "Xs wins again"
    fmt.Println(xoReferee([]string{"O.O ",
                                   "OO.",
                                   "XXX"})) //"X", "hp check"
}
