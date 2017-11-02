package main

import "fmt"

func ret(i int) int {
	o := 2
	return i * o
}

func main() {
	fmt.Println(ret(3))
}
