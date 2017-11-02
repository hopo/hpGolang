package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(romanNumerals(100))  // "hp Ex"
	fmt.Println(romanNumerals(6))    // 'VI', "First"
	fmt.Println(romanNumerals(76))   // 'LXXVI', "Second"
	fmt.Println(romanNumerals(499))  // 'CDXCIX', "Third"
	fmt.Println(romanNumerals(3888)) // 'MMMDCCCLXXXVIII', "Forth"
}

var rOne []string = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
var rTen []string = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
var rHun []string = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
var rTho []string = []string{"", "M", "MM", "MMM", "Mv", "v", "vM", "vMM", "vMMM", "Mx"}

func romanNumerals(number int) (ret string) {
	var s string = strconv.Itoa(number)
	var ss []string = strings.Split(s, "")
	var c0, c1, c2, c3 int

	switch len(ss) {
	case 1:
		c0, _ = strconv.Atoi(ss[0])
		ret = rOne[c0]
	case 2:
		c0, _ = strconv.Atoi(ss[0])
		c1, _ = strconv.Atoi(ss[1])
		ret = rTen[c0] + rOne[c1]
	case 3:
		c0, _ = strconv.Atoi(ss[0])
		c1, _ = strconv.Atoi(ss[1])
		c2, _ = strconv.Atoi(ss[2])
		ret = rHun[c0] + rTen[c1] + rOne[c2]
	case 4:
		c0, _ = strconv.Atoi(ss[0])
		c1, _ = strconv.Atoi(ss[1])
		c2, _ = strconv.Atoi(ss[2])
		c3, _ = strconv.Atoi(ss[3])
		ret = rTho[c0] + rHun[c1] + rTen[c2] + rOne[c3]
	}
	return ret
}
