package main

import (
	"fmt"
)

func main() {
	ex1 := bestSet(3, 13) // [4 4 5]
	ex2 := bestSet(2, 11) // [5 6]
	fmt.Println(ex1)
	fmt.Println(ex2)
}

func bestSet(n, s int) []int {
	return proof(n, s)
}

func proof(n, s int) []int {
	visl := vari_isl(n, s)
	var ret []int
	var max int
	m := 1
	for _, v := range visl {
		for _, w := range v {
			m *= w
		}
		if max < m {
			max = m
			ret = v
		}
		// fmt.Printf("v : %v\t m : %v\t max : %v\n", v, m, max) // chk all value
		m = 1
	}
	return ret
}

func vari_isl(n, s int) [][]int {
	if n == 2 {
		return vari_isl_two(n, s)
	}
	var ret, bc [][]int
	var a int
	for i := 1; i < s; i++ {
		a = s - i
		bc = vari_isl_two(2, i)
		for _, v := range bc {
			v = append(v, a)
			ret = append(ret, v)
		}
	}
	return ret
}

func vari_isl_two(n, s int) [][]int {
	var ret [][]int
	var a, b int
	for i := 1; i < s; i++ {
		a = s - i
		b = i
		ret = append(ret, []int{b, a})
	}
	return ret
}
