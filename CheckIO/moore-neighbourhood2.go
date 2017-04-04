package main

import "fmt"

if r == 0 { a,b,c = 0 }
if r == len(data)-1 {x,y,z = 0}

a : data[rm][cm]
b : data[rm][co]
c : data[rm][cp]
m : data[ro][cm]
n : data[ro][cm]
x : data[rp][co]
y : data[rp][co]
z : data[rp][cp]

func countNeighbours(data [][]int, r int, c int) int {
    var rmcm, rmco, rmcp, rocm, rocp, rpcm, rpco, rpcp int
    rm := r-1
    rp := r+1
    cm := c-1
    cp := c+1
    var rslt int

    // []int{data[rm][cm], data[rm][co], data[rm][cp], data[ro][cm], data[ro][cm], data[rp][co], data[rp][co], data[rp][cp]}

    checker := []int{data[rm][cm], data[rm][co], data[rm][cp], data[ro][cm], data[ro][cm], data[rp][co], data[rp][co], data[rp][cp]}

    for i := 0; i < len(checker); i++ {
        if checker[i] !=
        rslt += checker[i]
    }

    if r == 0 {
        rpco = data[r+1][c]
        if c == 0 {
            rocp = data[r][c+1]
            rpcp = data[r+1][c+1]
        } else if c == len(data[r])-1 {
            rocm = data[r][c-1]
            rpcm = data[r+1][c-1]
        } else {
            rocm = data[r][c-1]
            rocp = data[r][c+1]
            rpcm = data[r+1][c-1]
            rpcp = data[r+1][c+1]
        }
    }

    if r == len(data)-1 {
        rmco = data[r-1][c]
        if c == 0 {
            rmcp = data[r-1][c+1]
            rocp = data[r][c+1]
        } else if c == len(data[r])-1 {
            rmcm = data[r-1][c-1]
            rocm = data[r][c-1]
        } else {
            rmcm = data[r-1][c-1]
            rmcp = data[r-1][c+1]
            rocm = data[r][c-1]
            rocp = data[r][c+1]
        }
    }

    if 0 < r && r < len(data)-1 {
        rmco = data[r-1][c]
        rpco = data[r+1][c]
        if c == 0 {
            rmcp = data[r-1][c+1]
            rocp = data[r][c+1]
            rpcp = data[r+1][c+1]
        } else if c == len(data[r])-1 {
            rmcm = data[r-1][c-1]
            rocm = data[r][c-1]
            rpcm = data[r+1][c-1]
        } else {
            rmcm = data[r-1][c-1]
            rmcp = data[r-1][c+1]
            rocm = data[r][c-1]
            rocp = data[r][c+1]
            rpcm = data[r+1][c-1]
            rpcp = data[r+1][c+1]
        }
    }
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
