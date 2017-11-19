// *** api lab

package main

import (
	"fmt"
)

func main() {
	//fmt.Println(binary_convert(78)) // 1001110
	make_vari([]int{0, 1, 2})
}

// input int n, convert (type string)binary
func binary_convert(n int) string {
	//fmt.Println("n int:", n) // checker
	//fmt.Printf("%b\n", n) // checker
	var tsl []int
	for i := 1; i < n+1; i *= 2 {
		tsl = append([]int{i}, tsl...) // push front
	}

	zno := make([]int, len(tsl))
	for i := 0; i < len(tsl); i++ {
		n = n - tsl[i]
		if n == 0 {
			zno[i] = 1
			break
		}
		if n > 0 {
			zno[i] = 1
		}
		if n < 0 {
			n = n + tsl[i]
		}
	}

	var r string
	for _, v := range zno {
		r += string(v + 48) // int make string
	}
	return r
}

// variety
func make_vari(isl []int) {
	fmt.Println("isl:", isl)
	for i, _ := range isl {
		fmt.Println(isl[i])
	}
}
