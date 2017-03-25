package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

func main() {
	fmt.Println(Sqrt(2))
}
