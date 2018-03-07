package main

import (
	"fmt"
)

func main() {
	ex1 := toWeirdCase("try hello world")               // "TrY HeLlO WoRlD"
	ex2 := toWeirdCase("coding is thinking not typing") // "CoDiNg Is ThInKiNg NoT TyPiNg"
	fmt.Println(ex1)
	fmt.Println(ex2)
}

func toWeirdCase(s string) string {
	slsl := split_by_blank(s)
	for _, v := range slsl {
		for i, _ := range v {
			if i%2 == 0 {
				v[i] = minus_32(v[i])
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
	smpl := []byte(s)
	var ind []int
	for i, _ := range smpl {
		if smpl[i] == 32 {
			ind = append(ind, i)
		}
	}

	var ret [][]byte
	var box []byte
	a, b := 0, ind[0]
	box = smpl[:b]
	ret = append(ret, box)
	for i := 0; i < len(ind); i++ {
		if i == len(ind)-1 {
			box = smpl[b+1:]
			ret = append(ret, box)
			break
		}
		a, b = ind[i], ind[i+1]
		box = smpl[a+1 : b]
		ret = append(ret, box)
	}
	return ret
}

func minus_32(b byte) byte {
	return b - 32
}
