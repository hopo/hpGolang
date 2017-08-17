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

func (mo Moore) countNeighbours() int {
	var iss [][]int = mo.moore
	var r, c int = mo.row, mo.colomn
	var rm, rp int = r - 1, r + 1
	var cm, cp int = c - 1, c + 1

	rmcm := checker(iss, rm, cm)
	rmco := checker(iss, rm, c)
	rmcp := checker(iss, rm, cp)

	rocm := checker(iss, r, cm)
	rocp := checker(iss, r, cp)

	rpcm := checker(iss, rp, cm)
	rpco := checker(iss, rp, c)
	rpcp := checker(iss, rp, cp)

	return rmcm + rmco + rmcp + rocm + rocp + rpcm + rpco + rpcp
}

func main() {
	/*
		m1 := Moore{
			[][]int{{1, 0, 0, 1, 0}, {0, 1, 0, 0, 0}, {0, 0, 1, 0, 1}, {1, 0, 0, 0, 0}, {0, 0, 1, 0, 0}},
			1,
			2,
		}
	*/
	m2 := Moore{
		[][]int{{1, 0, 0, 1, 0}, {0, 1, 0, 0, 0}, {0, 0, 1, 0, 1}, {1, 0, 0, 0, 0}, {0, 0, 1, 0, 0}},
		0,
		0,
	}
	/*
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
	*/
	//fmt.Println(m1.countNeighbours()) //3, "1st example"
	fmt.Println(m2.countNeighbours()) //1, "2nd example"
	//fmt.Println(m3.countNeighbours()) //3, "Dense corner"
	//fmt.Println(m4.countNeighbours()) //0, "Single"
}
