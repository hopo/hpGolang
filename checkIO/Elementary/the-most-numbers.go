package main

import (
    "fmt"
    // "math"
    "sort"
)

func mostNumbers(nums ...float64) {
    if len(nums) == 0 {
        fmt.Println(0)
    } else {
        sort.Float64s(nums)

        fmt.Println(nums[len(nums)-1]-nums[0])
    }
}

func main() {
	mostNumbers(1, 2, 3)	//2, "3-1=2"
	mostNumbers(5, -5)	//10, "5-(-5)=10"
    mostNumbers(10.2, -2.2, 0, 1.1, 0.5)
	// mostNumbers(math.Ceil(mostNumbers(10.2, -2.2, 0, 1.1, 0.5)*1000))	//12400, "10.2-(-2.2)=12.4")
	mostNumbers() //0
}

/*
function mostNumbers(numbers){
	if(!numbers){return 0};
	var arr = Array.from(arguments).sort(function(a, b){return a-b});
	return arr[arr.length-1]-arr[0];
};
*/
