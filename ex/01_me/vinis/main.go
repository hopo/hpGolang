package main

import (
	"fmt"
	"spila"
)

func main() {
	is := []int{2, 2, 2, 6, 4}
	ex := spila.Vinis(0, is)
	if ex != nil {
		for _, v := range ex {
			fmt.Println(is[v])
		}
		return
	}
	fmt.Println(nil)
}
