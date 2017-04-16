package main

import (
	"fmt"
)

func countingTiles(radius float64) []float64 {
	return []float64{0, 0}
}

func main() {
	fmt.Println(countingTiles(2))   //[4, 12], "N=2"
	fmt.Println(countingTiles(3))   //[16, 20], "N=3"
	fmt.Println(countingTiles(2.1)) //[4, 20], "N=2.1"
	fmt.Println(countingTiles(2.5)) //[12, 20], "N=2.5"
}
