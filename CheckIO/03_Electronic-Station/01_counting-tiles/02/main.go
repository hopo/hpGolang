package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(countingTiles(2))          //[4, 12], "N=2"
	fmt.Println(countingTiles(3))          //[16, 20], "N=3"
	fmt.Println(countingTiles(2.1))        //[4, 20], "N=2.1"
	fmt.Println(countingTiles(2.5))        //[12, 20], "N=2.5"
	fmt.Println(countingTiles(math.Sqrt2)) // "hp ex"
}

type Tiles struct {
	solid   int
	partial int
}

var rt2 float64 = math.Sqrt2
var rt5 float64 = math.Sqrt(5)

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
	var ds, dr int

	switch {
	case n >= rt2*2:
		ds = 2
	case n >= rt2:
		ds = 1
	case n > 0:
		ds = 0
	}

	if n >= rt5 {
		dr = 2
	}

	ea = ds + dr
	return ea * 4
}

func pTiles(n float64) int {
	var ea int

	if n <= 3 {
		ea = 5
	}
	if n <= 2 {
		ea = 3
	}
	if n <= rt2 {
		ea = 2
	}

	return ea * 4
}
