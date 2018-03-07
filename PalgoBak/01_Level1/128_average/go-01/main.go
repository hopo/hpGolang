package main

import (
	"fmt"
	"github.com/hopo/hpkg"
)

func main() {
	ex1 := average([]float64{5, 3, 4}) // 4, 12/3
	fmt.Println(ex1)
}

func average(fs []float64) float64 {
	a := hpkg.Sum(fs)
	rslt := a / float64(len(fs))
	return rslt
}
