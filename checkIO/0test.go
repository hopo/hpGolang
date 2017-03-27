package main

import (
	//"strconv"
	"strings"
)

func main() {
	// var num int
	// num = 12045

	var str string
	// str = strconv.Itoa(num)
	str = "asd,bb"
	rslt := strings.split(str, ",")
	println(rslt)
}
