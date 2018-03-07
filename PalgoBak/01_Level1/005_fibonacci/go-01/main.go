package main

import (
	"fmt"
	"github.com/hopo/hpkg"
)

func main() {
	ex := fibonacci(3)
	fmt.Println(ex)
}

func fibonacci(num int) (ret int) {
	ret = hpkg.Fibo(num + 1)[num]
	return
}
