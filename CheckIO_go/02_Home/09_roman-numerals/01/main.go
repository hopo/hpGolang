package main

import (
	"fmt"
	"strconv"
)

var rNum map[byte]string = map[byte]string{
	'0': "",
	'1': "I",
	'2': "II",
	'3': "III",
	'4': "IV",
	'5': "V",
	'6': "VI",
	'7': "VII",
	'8': "VIII",
	'9': "IX",
}
var rNum1 map[byte]string = map[byte]string{
	'0': "",
	'1': "X",
	'2': "XX",
	'3': "XXX",
	'4': "XL",
	'5': "L",
	'6': "LX",
	'7': "LXX",
	'8': "LXXX",
	'9': "XC",
}
var rNum2 map[byte]string = map[byte]string{
	'0': "",
	'1': "C",
	'2': "CC",
	'3': "CCC",
	'4': "CD",
	'5': "D",
	'6': "DC",
	'7': "DCC",
	'8': "DCCC",
	'9': "CM",
}
var rNum3 map[byte]string = map[byte]string{
	'0': "",
	'1': "M",
	'2': "MM",
	'3': "MMM",
	'4': "IV",
	'5': "V",
	'6': "VI",
	'7': "VII",
	'8': "VIII",
	'9': "IX",
}

func reverseBs(bs []byte) []byte {
	var rbs []byte
	for i := len(bs) - 1; i > -1; i-- {
		rbs = append(rbs, bs[i])
	}
	return rbs
}

func romanNumerals(number int) (ret string) {
	var s string = strconv.Itoa(number)
	var bs []byte = []byte(s)
	var rbs []byte = reverseBs(bs)

	for i, v := range rbs {
		switch i {
		case 0:
			ret = rNum[v]
		case 1:
			ret = rNum1[v] + ret
		case 2:
			ret = rNum2[v] + ret
		case 3:
			ret = rNum3[v] + ret
		}
	}

	return ret
}

func main() {
	fmt.Println(romanNumerals(100))  // "hp Ex"
	fmt.Println(romanNumerals(6))    // 'VI', "First"
	fmt.Println(romanNumerals(76))   // 'LXXVI', "Second"
	fmt.Println(romanNumerals(499))  // 'CDXCIX', "Third"
	fmt.Println(romanNumerals(3888)) // 'MMMDCCCLXXXVIII', "Forth"
}
