package main

import (
	"fmt"
	"sort"
)

func mostNumbers(nums ...float64) float64 {
	if nums != nil {
		sort.Float64s(nums)
		res := nums[len(nums)-1] - nums[0]
		return res
	}
	return 0
}

func main() {
	fmt.Println(mostNumbers(1, 2, 3))                 //2, "3-1=2"
	fmt.Println(mostNumbers(5, -5))                   //10, "5-(-5)=10"
	fmt.Println(mostNumbers(10.2, -2.2, 0, 1.1, 0.5)) //12.4 "10.2-(-2.2)=12.4"
	fmt.Println(mostNumbers())                        //0 "Empty"
}

/*
function mostNumbers(numbers){
	if(!numbers){return 0};
	var arr = Array.from(arguments).sort(function(a, b){return a-b});
	return arr[arr.length-1]-arr[0];
};
*/
