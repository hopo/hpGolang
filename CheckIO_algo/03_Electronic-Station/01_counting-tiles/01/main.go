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

func countingTiles(radius float64) Tiles {
	var t Tiles

	t = Tiles{
		solid:   sTiles(radius),
		partial: pTiles(radius),
	}

	return t
}

func sTiles(n float64) int {
	var rt2 float64 = math.Sqrt2
	var rt5 float64 = math.Sqrt(5)

	var z0, z1 int

	for i := 1; i < 4; i++ {
		if n >= float64(i)*rt2 {
			z0 = i
		}
		if n >= float64(i)*rt5 {
			z1 = i
		}
	}
	return (z0 + (z1 * 2)) * 4
}

func pTiles(n float64) int {
	var cn float64 = math.Ceil(n)
	var x float64 = cn - n

	var rect float64 = math.Pow(cn*2, 2)
	var st int = sTiles(n)

	var xd float64 = x * 4
	var z int

	switch {
	case xd > 3:
		z = 3
	case xd > 1:
		z = 1
	}

	var pt int = int(rect) - st - (z * 4)

	return pt
}
