package main

import (
	"fmt"
	"strconv"
)

func main() {
	ex1 := no_continous("133303") // [1 3 0 3]
	ex2 := no_continous("47330")  // [4 7 3 0]
	fmt.Println(ex1)
	fmt.Println(ex2)
}

func no_continous(s string) (ret []int) {
	for i, v := range s {
		n, e := strconv.Atoi(string(v))
		if e != nil {
			panic(e)
		}
		if i != 0 && s[i-1] == s[i] {
			continue
		} else {
			ret = append(ret, n)
		}
	}
	return
}
