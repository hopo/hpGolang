package main

import (
	"fmt"
)

func main() {
	ex1 := expressions(15) // 4, [[1 2 3 4 5] [4 5 6] [7 8] [15]]
	ex2 := expressions(14) // 2, [[2 3 4 5] [14]]
	fmt.Println(ex1)
	fmt.Println(ex2)
}

func expressions(num int) interface{} {
	var chki int
	var box []int
	var ret [][]int
	for c := 0; c < num; c++ {
		for i := c + 1; i < num+1; i++ {
			chki += i
			box = append(box, i)
			if chki == num {
				if !member_in_islsl(ret, box) {
					ret = append(ret, box)
				}
				box = []int{}
				chki = 0
			}
			if chki > num {
				box = []int{}
				chki = 0
			}
		}
	}

	ret = sort_islsl(ret)
	return len(ret)
}

// sort islsl by 0 index
func sort_islsl(islsl [][]int) [][]int {
	ret := islsl
	for i, v := range ret {
		for j, w := range ret {
			if i != j && v[0] < w[0] {
				ret[i], ret[j] = ret[j], ret[i]
			}
		}
	}
	return ret
}

// member(isl) in islsl true or false
func member_in_islsl(s [][]int, m []int) bool {
	var chk int
	for _, v := range s {
		for i, w := range v {
			if w != m[i] {
				chk = 0
				break
			} else if w == m[i] {
				chk++
				if chk == len(m) {
					return true
				}
			}
		}
	}
	return false
}
