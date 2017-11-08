package main

import (
	"fmt"
)

func main() {
	ex := strange_sort([]string{"sun", "bed", "car"}, 1) // [car bed sun]
	fmt.Println(ex)
}

const ALP = "abcdefghijklmnopqrstuvwxyz"

func strange_sort(ssl []string, n int) (ret []string) {
	for _, r := range ALP {
		for _, v := range ssl {
			if byte(r) == v[n] {
				ret = append(ret, v)
			}
		}
	}
	return
}
