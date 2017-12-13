package main

import (
	"fmt"
	"github.com/hopo/hpkg"
)

func main() {
	ex1 := sumDivisor(12) // 28
	ex2 := sumDivisor(13) // 14
	fmt.Println(ex1)
	fmt.Println(ex2)
}

func sumDivisor(i int) (ret int) {
	d := hpkg.Denominator(i)
	ret = hpkg.Sum_isl(d)
	return
}
