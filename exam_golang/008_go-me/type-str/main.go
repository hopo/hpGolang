package main

import (
	"fmt"
	"strings"
)

type Str string

func (s Str) abc() {
	fmt.Println(s)

	xs := strings.Fields(string(s))
	for i, v := range xs {
		fmt.Println(i, ":", v)
	}
}

func main() {
	var s Str = "one little, two little, three little indian."
	s.abc()
}
