package main

import (
	"fmt"
)

func main() {
	//ex1 := jumpCase(4) // 5
	ex2 := jumpCase(5)
	//fmt.Println(ex1)
	fmt.Println(ex2)
}

func jumpCase(n int) int {
	fmt.Println("n:", n)
	fmt.Println("fibo:", fibo(n))

	dd := dduim(n)
	fmt.Println(dd)

	var slsl [][]int
	for _, v := range dd {
		if v == 2 {
			dd = br2(dd)
			slsl = append(slsl, dd)
		}
	}

	fmt.Println(slsl)
	return -1
}

func dduim(n int) []int {
	var isl []int
	var x int
	for {
		isl = append(isl, 2)
		x = n - 2
		n = x
		if x == 0 {
			break
		}
		if x == 1 {
			isl = append(isl, 1)
			break
		}
	}
	return isl
}

func br2(isl []int) []int {
	var ret []int
	for i, _ := range isl {
		j := len(isl) - 1 - i
		if isl[j] == 2 {
			ret = append(ret, isl[:j]...)
			ret = append(ret, []int{1, 1}...)
			ret = append(ret, isl[j+1:]...)
			break
		}
	}
	return ret
}

func fibo(n int) int {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return b
}
