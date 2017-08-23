package main

import (
	"fmt"
	"math"
)

func breakRings(rings [][]int) float64 {
	/*
	    //for understand
	    var r int
		for _, R := range rings {
			if r < R[0] {r = R[0]}
			if r < R[1] {r = R[1]}
		}
	    fmt.Printf("ring:%v, link:%v /", r, len(rings))
	*/
	return math.Ceil(float64(len(rings)) / 2)
}

func main() {
	fmt.Println(breakRings([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {4, 6}}))                         //3, "Example"
	fmt.Println(breakRings([][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}))                         //3, "All to all"
	fmt.Println(breakRings([][]int{{5, 6}, {4, 5}, {3, 4}, {3, 2}, {2, 1}, {1, 6}}))                         //3, "Chain"
	fmt.Println(breakRings([][]int{{8, 9}, {1, 9}, {1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}, {8, 7}})) //5, "Long chain"
}
