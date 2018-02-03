package main

import (
	"fmt"
)

func main() {
	ex1 := toWeirdCase("try hello world") // "TrY HeLlO WoRlD"
	//ex2 := toWeirdCase("coding is thinking not typing") // "CoDiNg Is ThInKiNg NoT TyPiNg"
	fmt.Println(ex1)
	//fmt.Println(ex2)
}

func toWeirdCase(s string) string {
	slsl := split_by_blank(s)
	for _, v := range slsl {
		for i, _ := range v {
			if i%2 == 0 {
				v[i] = minus_32(v[i]) // lower -> upper convert
			}
		}
	}
	var ret string
	for i, _ := range slsl {
		if i == len(slsl)-1 {
			ret += string(slsl[i])
		} else {
			ret += string(slsl[i]) + string(byte(32))
		}
	}
	return ret
}

func split_by_blank(s string) [][]byte {
	sample := []byte(s)
	var ret [][]byte
	var box []byte
	for i, v := range sample {
		if v == 32 {
			ret = append(ret, box)
			box = []byte{}
		} else {
			box = append(box, v)
		}
		if i == len(sample)-1 {
			ret = append(ret, box)
		}
	}
	return ret
}

func minus_32(b byte) byte {
	return b - 32
}
