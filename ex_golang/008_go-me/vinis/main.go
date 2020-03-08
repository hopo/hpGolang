package main

import (
	"fmt"
	"github.com/hopo/hpkg"
)

func main() {
	is := []int{2, 2, 2, 6, 4}
	ex := hpkg.Vinis(6, is)
	if ex != nil {
		for _, v := range ex {
			fmt.Println(is[v])
		}
		return
	}
	fmt.Println(nil)
}
