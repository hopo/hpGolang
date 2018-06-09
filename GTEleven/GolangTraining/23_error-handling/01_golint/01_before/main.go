package main

import "fmt"

func main() {
	x := 1
	str := evalInt(x)
	fmt.Println(str)
}

func evalInt(n int) string {
	if n > 10 {
		return fmt.Sprint("x is greater that 10")
	} else {
		return fmt.Sprint("x is less than 10")
	}
}
