package main

import (
    "fmt"
    "sort"
)

func main() {
    checkUnq([]int{10, 9, 10, 10, 9, 8})
}

func checkUnq(data []int) {
    x := data[0]

    i := sort.Search(len(data), func(i int) bool { return data[i] >= x })
    j := sort.Search(len(data), func(j int) bool { return data[j] <= x })
    if i != j {
        fmt.Println(data[i])
    }
}
