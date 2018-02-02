package main

import (
	"fmt"
)

func main() {
	ex1 := is_pair("(hello)()")                // true
	ex2 := is_pair(")(")                       // false
	ex3 := is_pair("(5+(40*(1+2)/2)+(1+2)*2)") // true
	fmt.Println(ex1)
	fmt.Println(ex2)
	fmt.Println(ex3)
}

func is_pair(s string) bool {
	bsl := []byte(s)
	var smpl []byte
	for _, v := range bsl {
		if v == 40 || v == 41 {
			smpl = append(smpl, v)
		}
	}
	l := len(smpl)
	var chk int
	for _, v := range smpl {
		if v == 40 {
			chk++
		}
	}
	a := l%2 == 0
	b := smpl[0] == 40
	c := smpl[l-1] == 41
	d := chk == l/2
	if a && b && c && d {
		return true
	}
	return false
}
