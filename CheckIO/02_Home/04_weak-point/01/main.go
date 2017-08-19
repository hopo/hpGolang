package main

import "fmt"

type Matrix [][]int

func (m Matrix) rowSum() []int {
	var rslt []int
	for _, v := range m {
		var s int
		for _, vv := range v {
			s += vv
		}
		rslt = append(rslt, s)
	}
	return rslt
}

func (m Matrix) columnSum() []int {
	var rslt []int
	for i, v := range m {
		var s int
		for j, _ := range v {
			s += m[j][i]
		}
		rslt = append(rslt, s)
	}
	return rslt
}

func lowerValueIndex(is []int) int {
	var c int = is[0]
	var rslt int
	for i, _ := range is {
		if is[i] < c {
			c = is[i]
			rslt = i
		}
	}
	return rslt
}

func (m Matrix) weakPoint() []int {
	rs := m.rowSum()
	cs := m.columnSum()

	lvi1 := lowerValueIndex(rs)
	lvi2 := lowerValueIndex(cs)

	rslt := []int{lvi1, lvi2}

	return rslt
}

func main() {
	m1 := Matrix{
		{7, 2, 7, 2, 8},
		{2, 9, 4, 1, 7},
		{3, 8, 6, 2, 4},
		{2, 5, 2, 9, 1},
		{6, 6, 5, 4, 5},
	}
	m2 := Matrix{
		{7, 2, 4, 2, 8},
		{2, 8, 1, 1, 7},
		{3, 8, 6, 2, 4},
		{2, 5, 2, 9, 1},
		{6, 6, 5, 4, 5},
	}
	m3 := Matrix{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}
	fmt.Println(
		m1.weakPoint(), // {3 3}, "Example"
		m2.weakPoint(), // {1 2}, "Two weak point"
		m3.weakPoint(), // {0 0}, "Top left"
	)
}
