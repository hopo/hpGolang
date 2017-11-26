package main

import (
	"fmt"
)

func main() {
	ex1 := first_word("Hello world") // "Hello"
	fmt.Println(ex1)
	ex2 := first_word(" a word ") // "a"
	fmt.Println(ex2)
	ex3 := first_word("don't touch it") // "don't"
	fmt.Println(ex3)
	ex4 := first_word("greetings, friends") // "greetings"
	fmt.Println(ex4)
	ex5 := first_word("... and so on ...") // "and"
	fmt.Println(ex5)
	ex6 := first_word("hi") // "hi"
	fmt.Println(ex6)
}

func first_word(text string) string {
	var ret []byte
	var uc, lc, us bool
	bsl := []byte(text)
	oi, ci := -1, -1

	for i, v := range bsl {
		uc = 64 < v && v < 91
		lc = 94 < v && v < 123
		if lc || uc {
			oi = i
			break
		}
	}

	for i, v := range bsl {
		uc = 64 < v && v < 91
		lc = 94 < v && v < 123
		us = v == 39
		if oi < i && !lc && !uc && !us {
			ci = i
			break
		}
	}

	if oi < 0 || ci < 0 {
		ret = bsl
	} else {
		ret = bsl[oi:ci]
	}
	return string(ret)
}
