package main

import (
	"fmt"
	"math"
)
func indexPower(sl []float64, n int) {
	var rslt float64
	if n < 0 {
		rslt = -1
	} else {
		tar := sl[n]
		fn := float64(n)
		rslt = math.Pow(tar, fn)
	}
	fmt.Println(rslt)
}

func main() {
	indexPower([]float64{1, 2, 3, 4}, 2)	//9
	indexPower([]float64{1, 3, 10, 100}, 3)	//1000000
	indexPower([]float64{0, 1}, 0)	//1
	indexPower([]float64{1, 2}, -1)	//-1
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
