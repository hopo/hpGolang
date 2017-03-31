package main

import (
	"fmt"
	"sort"
)

func median(data []float64) {
	sort.Float64s(data)
	fmt.Println(data)
}

func main() {
    median([]float64{1, 2, 3, 4, 5})	//3, "1st example"
    median([]float64{3, 1, 2, 5, 3})	//3, "2nd example"
	median([]float64{1, 300, 2, 200, 1})	//2, "3rd example"
    median([]float64{3, 6, 20, 99, 10, 15})	//12.5, "4th example"
}

/*
function median(data) {
	var sdata = data.sort(function(a, b) {return a-b})
	var x = data.length/2
	if(sdata.length%2 == 0) {
		return (sdata[x-1]+sdata[x])/2
	} else {
	    return sdata[Math.floor(x)]
	}
}
*/

