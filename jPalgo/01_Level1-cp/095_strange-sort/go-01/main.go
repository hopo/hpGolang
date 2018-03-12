package main

import (
	"fmt"
)

func main() {
	ex := strange_sort([]string{"sun", "bed", "car"}, 1) // ["car" "bed" "sun"]
	fmt.Println(ex)
}

func strange_sort(ssl []string, n int) (ret []string) {
	m := make(map[byte]string)
	for i, v := range ssl {
		m[v[n]] = ssl[i]
	}
	for i := 0; i < 200; i++ {
		for k, v := range m {
			if i == int(k) {
				fmt.Println(int(k))
				ret = append(ret, v)
			}
		}
	}
	fmt.Println("m:", m)
	return
}
