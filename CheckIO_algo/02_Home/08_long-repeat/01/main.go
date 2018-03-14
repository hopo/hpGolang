package main

import (
	"fmt"
	"strings"
)

func longRepeat(line string) int {
	ss := strings.Split(line, "")
	cf := ss[0]
	ea := 1
	ret := ea

	for i := 1; i < len(ss); i++ {
		if cf == ss[i] {
			ea++
		} else {
			cf = ss[i]
			ea = 1
		}
		if ret < ea {
			ret = ea
		}
	}
	return ret
}

func main() {
	s1 := "sdsffffse"
	s2 := "ddvvrwwwrggg"
	s3 := "ddvvvvvvvrwww"

	fmt.Println(longRepeat(s1)) // 4, "First"
	fmt.Println(longRepeat(s2)) // 3, "Second"
	fmt.Println(longRepeat(s3)) // 7, "hp ex"
}
