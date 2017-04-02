package main

import (
	"fmt"
	"math"
)
func indexPower(sl []float64, n int) float64 {
	if n < 0 { return -1 }
	tar, fn := sl[n], float64(n)
    return math.Pow(tar, fn)
}

func main() {
	fmt.Println(indexPower([]float64{1, 2, 3, 4}, 2))	//9
	fmt.Println(indexPower([]float64{1, 3, 10, 100}, 3))	//1000000
	fmt.Println(indexPower([]float64{0, 1}, 0))	//1
	fmt.Println(indexPower([]float64{1, 2}, -1))	//-1
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
