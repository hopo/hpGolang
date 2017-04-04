package main

import (
	"fmt"
    // "time"
)

type Data struct {
    rmcm, rmco, rmcp, rocm, rocp, rpcm, rpco, rpcp int
}

func main() {
    d := Data{}
    fmt.Println(d)
}

// func countNeighbours(d Data, r int, c int) {
//     fmt.Println(d.data[0], r, c)
    // for i := 0; i < len(checker); i++ {
        // rslt += checker[i]
    // }
// }

/*
func switchcase() {
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
*/

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
