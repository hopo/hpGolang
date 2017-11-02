package main

import (
	"fmt"
)

func main() {
	ex1 := printTriangle(3)
	fmt.Println(ex1)
}

func printTriangle(num int) string {
	n := "\n"
	rslt := fmt.Sprint("*", n, "**", n, "***")
	return rslt
}
