package main

import "fmt"

func zero(z *int) {
	fmt.Println("func zero :", z)
	*z = 0
}

func main() {
	x := 5
	fmt.Println(&x)
	zero(&x)
	fmt.Println(x) // x is 0
}
