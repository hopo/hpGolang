package main

import (
	"fmt"
	"math"
)

// I comment
type I interface {
	M()
}

// T comment
type T struct {
	S string
}

// T.M comment
func (t *T) M() {
	fmt.Println(t.S)
}

// F comment
type F float64

// F.M comment
func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
