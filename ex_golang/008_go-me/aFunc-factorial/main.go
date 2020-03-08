package main

import (
	"fmt"
)

func main() {
	ex1 := Fac01(5)
	ex2 := Fac02(5)
	ex3 := Fac03(5)
	ex4 := Fac04(5)
	ex5 := Fac05(5)
	fmt.Println(ex1, ex2, ex3, ex4, ex5)
}

func Fac01(x int) int {
	if x == 0 {
		return 1
	}
	return x * Fac01(x-1)
}

func Fac02(x int) int {
	for i := x; i > 2; i-- {
		x *= i - 1
	}
	return x
}

func Fac03(x int) int {
	ret := 1
	for i := 2; i < x+1; i++ {
		ret *= i
	}
	return ret
}

func Fac04(x int) int {
	for i := x; i > 2; i-- {
		x = func(i int) int {
			return x * (i - 1)
		}(i)
	}
	return x
}

func Fac05(x int) int {
	gen := func(i int) int {
		for i := x; i > 2; i-- {
			x *= (i - 1)
		}
		return x
	}
	return gen(x)
}
