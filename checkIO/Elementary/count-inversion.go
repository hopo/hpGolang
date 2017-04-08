package main

import "fmt"

func countInversion(sli []int) int {
	var check int
	for i := 0; i < len(sli); i++ {
		for j := 1; j < len(sli)-i; j++ {
			if sli[i] > sli[i+j] {
				check += 1
			}
		}
	}
    return check
}

func main() {
    fmt.Println(countInversion([]int{1, 2, 5, 3, 4, 7, 6})) //3
    fmt.Println(countInversion([]int{0, 1, 2, 3})) //0
    fmt.Println(countInversion([]int{99, -99})) //1
    fmt.Println(countInversion([]int{5, 3, 2, 1, 0})) //10
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
