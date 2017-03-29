package main

import "fmt"

func countInversion(sli []int) {
    fmt.Println(sli)
}
func main() {
    countInversion([]int{1, 2, 5, 3, 4, 7, 6}) //3
    countInversion([]int{0, 1, 2, 3}) //0
    countInversion([]int{99, -99}) //1
    countInversion([]int{5, 3, 2, 1, 0}) //10
}

/*
function countInversion(sequence){
	var check = 0;
	for(var i=0; i<sequence.length; i++){
		for(var n=1; n<sequence.length-i; n++){
			if(sequence[i]>sequence[i+n]){
				check += 1;
			};
		};
	};
	return check;
};
*/
