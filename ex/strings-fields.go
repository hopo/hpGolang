package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "string fields is string3 string4 string5 awesome"
	xs := strings.Fields(s)
	fmt.Println(xs[3])

	for i, v := range xs {
		fmt.Println(i, "-", v)
	}
}
