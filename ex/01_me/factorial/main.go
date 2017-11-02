package main

import (
	"fmt"
	"spila"
)

/*
func fac(l int) (r int) {
	r = 1
	for {
		r *= l * (l - 1)
		l = l - 1
		if l == 1 {
			return
		}
	}
	return -1
}
*/

/*
func fac(f int) (rslt int) {
	is := spila.Pyrange(1, f+1)
	rslt = 1
	for _, v := range is {
		rslt = rslt * v
	}
	return
}
*/

func main() {
	ex := spila.Fac(5)
	fmt.Println(ex)
}
