// https://tour.golang.org/flowcontrol/8
package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

func main() {
    var z float64
    for z = 0; z < 10 ; z++ {
        fmt.Println(z, ":", Sqrt(z))
    }
}
