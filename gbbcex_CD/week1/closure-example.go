package main

import "fmt"

func makeOddGenerator() func() int {
	i := int(1)
	return func() (r int) {
		r = i
		i += 2
		return
	}
}

func main() {
	nextOdd := makeOddGenerator()
	fmt.Println(nextOdd()) // 1
	fmt.Println(nextOdd()) // 3
	fmt.Println(nextOdd()) // 5
}
