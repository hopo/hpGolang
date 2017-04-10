package main

import (
	"fmt"
	"math"
	"sort"
)

// ByAbs struct
type ByAbs []float64

func (p ByAbs) Len() int           { return len(p) }
func (p ByAbs) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p ByAbs) Less(i, j int) bool { return math.Abs(p[i]) < math.Abs(p[j]) }

func absoluteSorting(nums []float64) []float64 {
	sort.Sort(ByAbs(nums))
	return nums
}

func main() {
	fmt.Println(
		absoluteSorting([]float64{-20, -5, 10, 15}), //[-5, 10, 15, -20]
		absoluteSorting([]float64{1, 2, 3, 0}),      //[0, 1, 2, 3]
		absoluteSorting([]float64{-1, -2, -3, 0}),   //[0, -1, -2, -3]
	)
}

/*
function absoluteSorting(numbers){
	return numbers.sort(function(a, b){return Math.abs(a)-Math.abs(b);});
}
*/
