package main

import (
	"fmt"
	"sort"
)

func median(data []float64) float64 {
	sort.Float64s(data)
	x := len(data)/2
	if len(data)%2 == 0 { return (data[x-1]+data[x])/2 }
	return data[x]
}

func main() {
    fmt.Println(median([]float64{1, 2, 3, 4, 5}))	//3, "1st example"
    fmt.Println(median([]float64{3, 1, 2, 5, 3}))	//3, "2nd example"
	fmt.Println(median([]float64{1, 300, 2, 200, 1}))	//2, "3rd example"
    fmt.Println(median([]float64{3, 6, 20, 99, 10, 15}))	//12.5, "4th example"
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
