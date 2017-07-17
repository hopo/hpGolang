package main

import "fmt"

func zero(x int) {
	x = 0
}

func zero2(y *int) {
	*y = 0
}

func main() {
	x := 5
	y := 5

	zero(x)
	zero2(&y)

	fmt.Println(x)
	fmt.Println(y)
}
