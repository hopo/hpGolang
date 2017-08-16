package main

import (
	"fmt"
	"math"
)

func indexPower(is []int, n int) int {
	if n < 0 {
		return n
	}
	tar, fn := is[n], n
	res := math.Pow(float64(tar), float64(fn))
	return int(res)
}

func main() {
	fmt.Println(indexPower([]int{1, 2, 3, 4}, 2))    //9
	fmt.Println(indexPower([]int{1, 3, 10, 100}, 3)) //1000000
	fmt.Println(indexPower([]int{0, 1}, 0))          //1
	fmt.Println(indexPower([]int{1, 2}, -1))         //-1
}

/*
function indexPower(array, n){
    if(array[n]){
		return Math.pow(array[n], n);
	}else if(array[n]==0){
		return 1;
	};
	return -1;
};
*/
