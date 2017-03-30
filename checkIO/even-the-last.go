package main

import (
	"fmt"
)

func evenLast(nums []int) {
	sum := 0
	for i:= 0; i < len(nums); i++ {
		if i % 2 == 0 {
			sum += nums[i]
		}
	}
	fmt.Println(sum*nums[len(nums)-1])
}

func main() {
    evenLast([]int{0, 1, 2, 3, 4, 5})	//30, "(0+2+4)*5=30"
    evenLast([]int{1, 3, 5})	//30, "(1+5)*5=30"
    evenLast([]int{6})	//36, "(6)*6=36"
    //evenLast([]int)	//0, "An empty array = 0"
}

/*
function evenLast(data) {
    if(data.length){
		var sum = 0;
		for(var n = 0; n < data.length; n = n + 1){
			if(n % 2 == 0){
				sum = sum + data[n]
			};
		};
		return sum * data[data.length-1]
	}else{
		return 0;
	};
};
*/
