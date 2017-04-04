package main

import (
	"fmt"
    // "time"
)


func main() {
    data := [][]int{
        {1,2,3},
        {1,2,3},
    }
    r, c := 0, 2

    switch r == 0{
    case c == 1:
        fmt.Println("c == 1")
    case c == 2:
        fmt.Println("c == 2")
    }
    fmt.Println(data)
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
