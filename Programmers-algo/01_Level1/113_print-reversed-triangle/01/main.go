package main

import (
	"fmt"
	"github.com/hopo/hpkg"
)

func main() {
	ex := printReversedTriangle(3)
	fmt.Println(ex)
}

func printReversedTriangle(n int) (ret string) {
	for i := 0; i < n; i++ {
		q := n - i
		r := hpkg.Smulti("*", q)
		switch q {
		default:
			ret += fmt.Sprintln(r)
		case 1:
			ret += fmt.Sprint(r)
		}
	}
	return
}
