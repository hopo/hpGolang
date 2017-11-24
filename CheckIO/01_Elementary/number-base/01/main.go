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
	return -1
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
