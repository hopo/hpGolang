package main

import (
	"fmt"
	"math"
)

type ByAbs []int

func (p ByAbs) Len() int {
	return len(p)
}

func (p ByAbs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p ByAbs) Less(i, j int) bool {
	return math.Abs(float64(p[i])) < math.Abs(float64(p[j]))
}

func (p ByAbs) absoluteSorting() []int {
	l := p.Len()
	for i := 1; i < l; i++ {
		for j := 0; j < l-1; j++ {
			if p.Less(i, j) {
				p.Swap(i, j)
			}
		}
	}
	return p
}

func main() {
	var p1, p2, p3 ByAbs
	p1 = []int{-20, -5, 10, 15}
	p2 = []int{1, 2, 3, 0}
	p3 = []int{-1, -2, -3, 0}

	fmt.Println(p1.absoluteSorting()) //[-5, 10, 15, -20]
	fmt.Println(p2.absoluteSorting()) //[0, 1, 2, 3]
	fmt.Println(p3.absoluteSorting()) //[0, -1, -2, -3]
}
