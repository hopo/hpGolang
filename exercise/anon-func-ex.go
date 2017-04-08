//example 1
func bar() int {
	var f func(int, int) int
	f = func(a int, b int) int {
		return a * b
	}
	return f(3, 5)
}

//example 2
func bar() int {
	f := func(a int, b int) int {
		return a * b
	}
	return f(3, 5)
}

//example 3
func bar() int {
	return func(a int, b int) int {
		return a * b
	}(3, 5)
}

//example 4
func bar() func(int, int) int {
	return func(a int, b int) int {
		return a * b
	}
}
