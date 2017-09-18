package main

import (
	"fmt"
	"math"
)

func main() {
	// code
	sqrt5ch()
}

func squareDiagonal(w float64) float64 {
	wPow := math.Pow(w, 2)
	d := math.Sqrt(wPow + wPow)
	return d
}

func sqrt2ch() {
	for i := 1; i < 4; i++ {
		fmt.Println("sqrt2 *", i, ":", math.Sqrt2*float64(i))
	}
}

func sqrt5ch() {
	for i := 1; i < 4; i++ {
		fmt.Println("sqrt5 *", i, ":", math.Sqrt(5)*float64(i))
	}
}

func conferRnC(num float64) {
	n := math.Ceil(num) * 2
	rect := n * n
	circ := math.Pow(num, 2) * math.Pi

	c := rect - circ

	fmt.Printf("rect: %v - circ: %v  = %v \n", rect, circ, c)
}
