package main

import "fmt"

type Moore struct {
	moore  [][]int
	row    int
	colomn int
}

func eRecover() {
	err := recover()
	if err != nil {
		return
	}
}

func checker(iss [][]int, n int, m int) int {
	defer eRecover()
	return iss[n][m]
}

func (m Moore) countNeighbours() int {
	var iss [][]int = m.moore
	var r, c int = m.row, m.colomn

	var res int
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			res += checker(iss, r-i, c-j)
		}
	}
	return res - iss[r][c]
}

func main() {
	m1 := Moore{
		[][]int{{1, 0, 0, 1, 0}, {0, 1, 0, 0, 0}, {0, 0, 1, 0, 1}, {1, 0, 0, 0, 0}, {0, 0, 1, 0, 0}},
		1,
		2,
	}
	m2 := Moore{
		[][]int{{1, 0, 0, 1, 0}, {0, 1, 0, 0, 0}, {0, 0, 1, 0, 1}, {1, 0, 0, 0, 0}, {0, 0, 1, 0, 0}},
		0,
		0,
	}
	m3 := Moore{
		[][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
		0,
		2,
	}
	m4 := Moore{
		[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
		1,
		1,
	}

	fmt.Println(m1.countNeighbours()) //3, "1st example"
	fmt.Println(m2.countNeighbours()) //1, "2nd example"
	fmt.Println(m3.countNeighbours()) //3, "Dense corner"
	fmt.Println(m4.countNeighbours()) //0, "Single"
}
