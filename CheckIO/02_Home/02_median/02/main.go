package main

import (
	"fmt"
	"sort"
)

func median(data []float64) float64 {
	sort.Float64s(data)
	l := len(data)
	if l%2 == 0 {
		return (data[(l/2)-1] + data[l/2]) / 2
	}
	return data[l/2]
}

func main() {
	fmt.Println(median([]float64{1, 2, 3, 4, 5}))        //3, "1st example"
	fmt.Println(median([]float64{3, 1, 2, 5, 3}))        //3, "2nd example"
	fmt.Println(median([]float64{1, 300, 2, 200, 1}))    //2, "3rd example"
	fmt.Println(median([]float64{3, 6, 20, 99, 10, 15})) //12.5, "4th example"
}
