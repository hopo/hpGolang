package main

import (
	"fmt"
)

func main() {
	fmt.Println(foo(7, 59))
}

func foo(i, n int) int {
	defer rcv()
	a := make([]int, 4)
	a[i] = n // panic: out of bounds!
	return a[i]
}

func rcv() {
	err := recover()
	if err != nil {
		fmt.Println("Recovered from this error:")
		fmt.Printf("*err* %v; %T\n", err, err)
	}
}
