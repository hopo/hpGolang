package main

import (
	"fmt"
	"sort"
)

func main() {
	ex1 := most_frequent([]string{
		"a", "b", "c",
		"a", "b",
		"a",
	}) // 'a'
	fmt.Println(ex1)
	ex2 := most_frequent([]string{"g", "a", "bi", "bi", "bi"}) // "bi"
	fmt.Println(ex2)
}

type Max struct {
	value string
	count int
}

func most_frequent(data []string) string {
	sort.Strings(data)
	var count int
	chkstr := data[0]
	m := Max{chkstr, 1}

	for _, v := range data {
		if chkstr == v {
			count++
		}
		if chkstr != v {
			if m.count < count {
				m.count = count
				m.value = chkstr
			}
			chkstr = v
			count = 1
		}
	}

	return m.value
}
