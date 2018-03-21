package main

import (
	"fmt"
)

func main() {
	ex1 := hide_numbers("01033334444") // "*******4444"
	ex2 := hide_numbers("027778888")   // "*****8888"
	fmt.Println(ex1)
	fmt.Println(ex2)
}

func hide_numbers(s string) (ret string) {
	sl := []byte(s)
	l := len(sl)
	for i, _ := range s {
		if l-i > 4 {
			sl[i] = '*'
		}
		ret += string(sl[i])
	}
	return
}
