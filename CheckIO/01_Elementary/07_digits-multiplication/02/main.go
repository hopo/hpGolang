package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func digitsMultip(data int) int {
	s := strconv.Itoa(data)
	cs := regexp.MustCompile("[0]").ReplaceAllString(s, "")
	res := 1
	var ss []string
	for i := 0; i < len(cs); i++ {
		ss = append(ss, string(cs[i]))
	}
	for _, v := range ss {
		val, _ := strconv.Atoi(v)
		res *= val
	}
	return res
}

func main() {
	fmt.Println(digitsMultip(123405)) //120
	fmt.Println(digitsMultip(999))    //729
	fmt.Println(digitsMultip(1000))   //1
	fmt.Println(digitsMultip(1111))   //1
}
