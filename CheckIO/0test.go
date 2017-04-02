package main

import (
	"fmt"
	// "sort"
)

func main() {
    data := []int{5, 6, 9}
    for i, _ := range data {
        if data[i] { fmt.Println(data[i]) } else { fmt.Println("-") }
    }

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
