package main

func indexPower([]int, int) int{}

func main() {
	indexPower([1, 2, 3, 4], 2)	//9
	indexPower([1, 3, 10, 100], 3)	//1000000
	indexPower([0, 1], 0)	//1
	indexPower([1, 2], -1)	//-1
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
