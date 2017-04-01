package main

import (
	"fmt"
	"sort"
)

func main() {
    data := []int{1, 9, 4, 5, 4, 2}
    i := data[2]
    sort.Ints(data)
    s := sort.SearchInts(data, i)
    // i := 4
    // data = append(data[:i], data[i+1:]...)
    fmt.Println(data, s)
}
