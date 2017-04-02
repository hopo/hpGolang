package main

import (
	"fmt"
	// "sort"
)

func main() {
    nonUniqueElements([]int{0, 1, 2, 1})

}

/*
func nonUniqueElements(data []int) {
    var rslt []int
    for i, d := range data {
        for j, _ := range data {
            if i != j && d == data[j] {
                rslt = append(rslt, d)
                break
            }
        }
    }
    fmt.Println(rslt)
}
*/
