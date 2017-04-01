package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{10, 9, 10, 10, 9, 8}
	x := 10

	i := sort.Search(len(a), func(i int) bool { return a[i] >= x })
	if i < len(a) && a[i] == x {
		fmt.Printf("found %d at index %d in %v\n", x, i, a)
	} else {
		fmt.Printf("%d not found in %v\n", x, a)
	}
}
