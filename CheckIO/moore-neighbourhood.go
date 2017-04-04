package main

import "fmt"

func checker(data [][]int, n int, m int) int {
    defer func() {
        err := recover()
        if err != nil { return }
    }()
    return data[n][m]
}

func countNeighbours(data [][]int, r int, c int) int {
    var rm, rp, cm, cp int = r-1, r+1, c-1, c+1
    rmcm := checker(data, rm, cm)
    rmco := checker(data, rm, c)
    rmcp := checker(data, rm, cp)

    rocm := checker(data, r, cm)
    rocp := checker(data, r, cp)

    rpcm := checker(data, rp, cm)
    rpco := checker(data, rp, c)
    rpcp := checker(data, rp, cp)


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

/*
function countNeighbours(data, row, col) {
  var topl, topc, topr, bdyl, bdyr, botl, botc, botr;
  if(row == 0) {
    topl = 0, topc = 0, topr = 0
  } else {
    topl = data[row-1][col-1] == undefined ? 0 : data[row-1][col-1];
    topc = data[row-1][col] == undefined ? 0 : data[row-1][col];
    topr = data[row-1][col+1] == undefined ? 0 : data[row-1][col+1];
  }

  bdyl = data[row][col-1] == undefined ? 0 : data[row][col-1];
  bdyr = data[row][col+1] == undefined ? 0 : data[row][col+1];

  if(row == data.length-1){
    botl = 0, botc = 0, botr = 0;
  } else {
    botl = data[row+1][col-1] == undefined ? 0 : data[row+1][col-1];
    botc = data[row+1][col] == undefined ? 0 : data[row+1][col];
    botr = data[row+1][col+1] == undefined ? 0 : data[row+1][col+1];
  }

  var rslt = topl+topc+topr+bdyl+bdyr+botl+botc+botr;
  return rslt;
}
*/
