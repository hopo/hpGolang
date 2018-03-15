package main

import (
	"fmt"
)

func main() {
	ex1 := longest_palindrom("tomotmomtomot") // 13
	ex2 := longest_palindrom("tomotmkdir")    // 5
	ex3 := longest_palindrom("xmmxaj")        // 4
	fmt.Println(ex1)
	fmt.Println(ex2)
	fmt.Println(ex3)
}

func longest_palindrom(s string) int {
	bsl := []byte(s)
	start := bsl[0]
	sample := sameSE(start, bsl)

	if mirrorChk(sample) {
		//fmt.Println(string(sample)) // show string
		return len(sample)
	}

	return -1
}

func sameSE(value uint8, bsl []byte) []byte {
	var ret []byte
	l := len(bsl)
	for i := l - 1; i > 0; i-- {
		if value == bsl[i] {
			ret = bsl[:i+1]
			break
		}
	}
	return ret
}

func mirrorChk(sample []byte) bool {
	l := len(sample)
	var fhalf, shalf []byte

	fhalf = sample[:l/2]
	if l%2 == 0 {
		shalf = sample[l/2:]
	} else {
		shalf = sample[l/2+1:]
	}

	for i := 0; i < len(fhalf); i++ {
		j := l/2 - 1 - i
		if fhalf[i] != shalf[j] {
			return false
		}
	}
	return true
}
