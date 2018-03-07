package main

import (
	"fmt"
	"github.com/hopo/hpkg"
)

func main() {
	ex := digit_reverse(12345) // [5 4 3 2 1]
	fmt.Println(ex)
}

func digit_reverse(n int) []int {
	isl := hpkg.Int_convert_each(n)
	var ret []int
	for i := 0; i < len(isl); i++ {
		j := len(isl) - 1 - i
		ret = append(ret, isl[j])
	}
	return ret
}
