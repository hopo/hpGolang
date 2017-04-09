package main

import "fmt"

//Nums struct
type Nums struct {
	a, b int
}

func (n Nums) sum() int {
	return n.a + n.b
}

func main() {
	nums := Nums{1, 2}
	sum := nums.sum()
	fmt.Println(sum)
}
