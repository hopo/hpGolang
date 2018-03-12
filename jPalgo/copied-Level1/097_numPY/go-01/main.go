package main

import (
	"fmt"
)

func main() {
	ex1 := numPY("pPoooyY") // true
	ex2 := numPY("Pyy")     // false
	fmt.Println(ex1)
	fmt.Println(ex2)
}

func numPY(s string) bool {
	var ppp, yyy int
	for _, v := range s {
		switch string(v) {
		case "p":
			ppp++
		case "y":
			yyy++
		}
	}
	return ppp == yyy
}
