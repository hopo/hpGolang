package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(countingTiles(2))            //[4, 12], "N=2"
	fmt.Println(countingTiles(3))            //[16, 20], "N=3"
	fmt.Println(countingTiles(2.1))          //[4, 20], "N=2.1"
	fmt.Println(countingTiles(2.5))          //[12, 20], "N=2.5"
	fmt.Println(countingTiles(math.Sqrt(5))) // "hp ex"
}

type Tiles struct {
	solid   int
	partial int
}

func countingTiles(radius float64) Tiles {
	var st int = sTiles(radius)
	var pt int = pTiles(radius)

	var t Tiles = Tiles{
		solid:   st,
		partial: pt,
	}

	return t
}

func sTiles(n float64) int {
	var ea int
	switch {
	case n < math.Sqrt2:
		ea = 0
	case n < math.Sqrt(5):
		ea = 1
	case n < math.Sqrt2*2:
		ea = 3
	case n >= math.Sqrt2*2:
		ea = 4
	}

	return ea * 4
}

func pTiles(n float64) int {
	var ea int
	switch {
	case n < math.Sqrt2:
		ea = 1
	case n == math.Sqrt2:
		ea = 2
	case n <= 2 || n == math.Sqrt(5):
		ea = 3
	case n <= 3:
		ea = 5
	}

	return ea * 4
}
