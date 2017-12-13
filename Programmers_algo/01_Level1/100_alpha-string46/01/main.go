package main

import (
	"fmt"
	"github.com/hopo/hpkg"
)

func main() {
	ex1 := alpha_string46("a234") // False, "string 'a'"
	ex2 := alpha_string46("1234") // True, "all int"

	fmt.Println(ex1)
	fmt.Println(ex2)
}

func alpha_string46(str string) bool {
	return hpkg.Isdigit(str)
}
