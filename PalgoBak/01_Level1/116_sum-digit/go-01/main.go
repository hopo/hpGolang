package main

import (
	"fmt"
)

func main() {
	ex1 := sum_digit(123)       // 6
	ex2 := sum_digit(999999999) // 81
	fmt.Println(ex1)
	fmt.Println(ex2)

}

func sum_digit(n int) (ret int) {
	di := 1
	for {
		if n-di < 0 {
			di = di / 10
			break
		}
		di *= 10
	}

	for {
		if di == 1 {
			ret += n
			break
		}
		z := n / di
		ret += z
		n = n % di
		di = di / 10
	}
	return ret
}
