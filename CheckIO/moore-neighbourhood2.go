package main

import "fmt"

type Data struct{
    sli [][]int
    r int
    c int
}

func (dt Data) checker() int {
    defer func() {
        err := recover()
        if err != nil { return }
    }()
    return dt.sli[dt.r][dt.c]
}

func countNeighbours(sli [][]int, r int, c int) int {
    // neig := []int{sli[r-1][c-1], sli[r-1][c], sli[r-1][c+1], sli[r][c-1], sli[r][c+1], sli[r+1][c-1], sli[r+1][c], sli[r+1][c+1]}
    rmcm := Data{sli, r-1, c-1}.checker()
    rmco := Data{sli, r-1, c}.checker()
    rmcp := Data{sli, r-1, c+1}.checker()

    rocm := Data{sli, r, c-1}.checker()
    rocp := Data{sli, r, c+1}.checker()

    rpcm := Data{sli, r+1, c-1}.checker()
    rpco := Data{sli, r+1, c}.checker()
    rpcp := Data{sli, r+1, c+1}.checker()


    return rmcm+rmco+rmcp+rocm+rocp+rpcm+rpco+rpcp
}


func main() {
    fmt.Println(countNeighbours([][]int{{1, 0, 0, 1, 0},
                                        {0, 1, 0, 0, 0},
                                        {0, 0, 1, 0, 1},
                                        {1, 0, 0, 0, 0},
                                        {0, 0, 1, 0, 0}}, 1, 2)) //3, "1st example"
    fmt.Println(countNeighbours([][]int{{1, 0, 0, 1, 0},
                                        {0, 1, 0, 0, 0},
                                        {0, 0, 1, 0, 1},
                                        {1, 0, 0, 0, 0},
                                        {0, 0, 1, 0, 0}}, 0, 0)) //1, "2nd example"
    fmt.Println(countNeighbours([][]int{{1, 1, 1},
                                        {1, 1, 1},
                                        {1, 1, 1}}, 0, 2))   //3, "Dense corner"
    fmt.Println(countNeighbours([][]int{{0, 0, 0},
                                        {0, 1, 0},
                                        {0, 0, 0}}, 1, 1))    // 0, "Single"
}
