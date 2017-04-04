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
    r, c := 1, 1

    switch r == 0{
    case c == 0:
        fmt.Println(data[-1][0])
    case c == 1:
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
