package main

import (
	"fmt"
	"strconv"
)

func main() {
	ex := strToInt("-1234")
	fmt.Println(ex)
}

func strToInt(s string) (ret int) {
	ret, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return
}
