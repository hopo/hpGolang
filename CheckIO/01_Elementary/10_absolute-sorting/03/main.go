package main

import (
	"fmt"
	"math"
	//"sort"
)

// ByAbs struct
type ByAbs []int

func (p ByAbs) Len() int           { return len(p) }
func (p ByAbs) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p ByAbs) Less(i, j int) bool { return math.Abs(float64(p[i])) < math.Abs(float64(p[j])) }

func absoluteSorting(is []int) []int {
	var p ByAbs = is

	fmt.Println(p.Less(p[1], p[0]))
	return p
}

func main() {
	fmt.Println(absoluteSorting([]int{-20, -5, 10, 15})) //[-5, 10, 15, -20]
	//fmt.Println(absoluteSorting([]int{1, 2, 3, 0}))      //[0, 1, 2, 3]
	//fmt.Println(absoluteSorting([]int{-1, -2, -3, 0}))   //[0, -1, -2, -3]
}

/*
function absoluteSorting(numbers){
	return numbers.sort(function(a, b){return Math.abs(a)-Math.abs(b);});
}
*/
