package main

import (
	"fmt"
)

func main() {
    ss := []int{1, 2, 3, 4, 5}

    p := ss[7]
    fmt.Println(p)
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
