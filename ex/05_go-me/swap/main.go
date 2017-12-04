package main

import (
	"fmt"
)

func main() {
	a, b := 2, 5
	fmt.Println(a, b)
	a, b = b, a
	fmt.Print(a, b)
}
