package main

import "fmt"

func main() {
	ex1 := bar01()
	ex2 := bar02()
	ex3 := bar03()
	ex4 := bar04()
	fmt.Println(ex1, ex2, ex3, ex4)
}

//example 1
func bar01() int {
	var f func(int, int) int
	f = func(a int, b int) int {
		return a * b
	}
	return f(3, 5)
}

//example 2
func bar02() int {
	f := func(a int, b int) int {
		return a * b
	}
	return f(3, 5)
}

//example 3
func bar03() int {
	return func(a int, b int) int {
		return a * b
	}(3, 5)
}

//example 4
func bar04() func(int, int) int {
	return func(a int, b int) int {
		return a * b
	}
}
