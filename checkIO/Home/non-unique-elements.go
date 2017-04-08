package main

import (
	"fmt"
)

func nonUniqueElements(data []int) []int {
    var rslt []int
    for i, d := range data {
        for j, _ := range data {
            if i != j && d == data[j] {
                rslt = append(rslt, d)
                break
            }
        }
    }
    return rslt
}

func main() {
    fmt.Println(nonUniqueElements([]int{1, 2, 3, 1, 3}))	//[1, 3, 1, 3], "1st example"
    fmt.Println(nonUniqueElements([]int{1, 2, 3, 4, 5}))	//[], "2nd example"
    fmt.Println(nonUniqueElements([]int{5, 5, 5, 5, 5}))	//[5, 5, 5, 5, 5], "3rd example"
    fmt.Println(nonUniqueElements([]int{10, 9, 10, 10, 9, 8}))	//[10, 9, 10, 10, 9], "4th example"
}

/*
function nonUniqueElements(data) {
	for(var i=0; i<data.length; i++) {
		var sock = data[i];
		delete data[i];
		data.includes(sock) ? data[i] = sock : null;
	};
	return data.filter(() => {return true});
}
*/

/*
nonUniqueElements = a => a.filter(e => a.indexOf(e) != a.lastIndexOf(e));
*/
