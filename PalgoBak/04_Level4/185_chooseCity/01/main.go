package main

import (
	"fmt"
)

func main() {
	ex1 := chooseCity(3, [][]int{{1, 5}, {2, 2}, {3, 3}})         // 1
	ex2 := chooseCity(4, [][]int{{1, 2}, {2, 3}, {3, 7}, {4, 3}}) // 3
	fmt.Println(ex1)
	fmt.Println(ex2)
}

func chooseCity(n int, city [][]int) interface{} {
	var report []int
	var chk, f int
	for i := 0; i < n; i++ {
		f = city[i][0]
		for j := 0; j < n; j++ {
			if i != j {
				chk += distance(f, city[j])
			}
		}
		report = append(report, chk)
		chk = 0
	}
	// fmt.Printf("report: %v\n", report) // checker
	min, midx := report[0], 0
	for i, v := range report {
		if min > v {
			min = v
			midx = i
		}
	}
	return midx + 1 // index begin 0
}

func distance(f int, to []int) int {
	d := to[0] - f
	d = absolute(d)
	return d * to[1]
}

func absolute(n int) int {
	if n < 0 {
		n *= -1
	}
	return n
}
