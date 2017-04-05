package main

import (
	"fmt"
    // "log"
    // "os"
    // "time"
)

/*
func chcker(data [][]int, r int, c int) int {
    defer func() {
        err := recover()
        if err != nil { return }
    }()
    return data[r][c]
}

func main() {
    d := chcker([][]int{{1, 0, 0}, {0, 1, 0}, {0, 1, 1}}, 0, 0)
    fmt.Println(d)
}
/*


/*
type MyError struct{
    what string
}

func (e MyError) Error() string {
    return "Fucking-Error"
}

func run() (bl bool, er error) {
    bl = 1 < 0
    er = MyError{ "a" }
    return
}
func main(){
    b, err := run()
    if err != nil {
        fmt.Println(b, err)
    }
}
*/

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
