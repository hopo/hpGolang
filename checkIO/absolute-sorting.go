package main

import (
    "fmt"
    // "math"
    "sort"
)

func absoluteSorting(nums []float64) {
    func(a, b float64) bool {
        return math.Abs(a) < math.Abs(b)
    }
    // sort.Float64s(nums)
    fmt.Println(nums)
}

func main() {
    absoluteSorting([]float64{-20, -5, 10, 15})  //[-5, 10, 15, -20]
    absoluteSorting([]float64{1, 2, 3, 0})   //[0, 1, 2, 3]
    absoluteSorting([]float64{-1, -2, -3, 0})    //[0, -1, -2, -3]
}

/*
function absoluteSorting(numbers){
	return numbers.sort(function(a, b){return Math.abs(a)-Math.abs(b);});
}
*/
