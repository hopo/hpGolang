package main

import (
	"fmt"
)

func main() {
	ex1 := number_base("AF", 16) // 175
	fmt.Println(ex1)
	ex2 := number_base("101", 5) // 26
	fmt.Println(ex2)
}

func number_base(s string, n int) int {
	bsl := []byte(s)
	var isl []int
	for _, v := range bsl {
		if 64 < v && v < 91 {
			isl = append(isl, int(v-55))
		}
		if 47 < v && v < 59 {
			isl = append(isl, int(v-48))
		}
	}

	pow := func(n int, p int) int {
		ret := 1
		if p == 0 {
			return ret
		}
		for i := 0; i < p; i++ {
			ret *= n
		}
		return ret
	}

	var j int
	var ret int
	for i, _ := range isl {
		j = len(isl) - 1 - i
		ret += isl[j] * pow(n, i)
	}
	return ret
}

/*
ALP = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'

def checkio(str_number, radix):
    l = len(str_number)
    il = []
    for i in range(l):
        if str_number[i] in ALP:
            e = ALP.index(str_number[i]) + 10
        else:
            e = int(str_number[i])
        if e >= radix:
            return -1
        else:
            r = e * radix**(l-(i+1))
            il.append(r)
    return sum(il)

if __name__ == '__main__':
    ex1 = checkio("AF", 16) # 175
    print(ex1)
    ex2 = checkio("101", 5) # 26
    print(ex2)
*/
