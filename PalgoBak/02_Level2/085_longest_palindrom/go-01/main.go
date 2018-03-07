package main

import (
	"fmt"
)

func main() {
	ex1 := longest_palindrom("tomotmomtomot") // 13
	ex2 := longest_palindrom("tomotmkdir")    // 5
	fmt.Println(ex1)
	fmt.Println(ex2)
}

func longest_palindrom(s string) int {
	bsl := []byte(s)
	start := bsl[0]
	li := findLastIndex(start, bsl)

	var tgsli []byte
	tgsli = append(tgsli, bsl[:li+1]...)

	if mirrorChk(tgsli) {
		//s = string(tgsli)
		//fmt.Println(s)
		return li + 1
	}

	return -1
}

func findLastIndex(value uint8, bsl []byte) int {
	l := len(bsl)
	for i := l - 1; i > 0; i-- {
		if value == bsl[i] {
			return i
		}
	}
	return -1
}

func mirrorChk(bsl []byte) bool {
	l := len(bsl)
	h := l / 2
	var slia, slib []byte
	slia = append(slia, bsl[:h]...)
	slib = append(slib, bsl[h+1:]...)
	for i := 0; i < h; i++ {
		if slia[i] == slib[h-1] {
			return true
		}
	}
	return false

}
