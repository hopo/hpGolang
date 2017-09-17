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

func reverseSs(ss []string) []string {
	var rss []string
	for i := len(ss) - 1; i > -1; i-- {
		rss = append(rss, ss[i])
	}
	return rss
}

func romanNumerals(number int) (ret string) {
	var s string = strconv.Itoa(number)
	var ss []string = strings.Split(s, "")
	var rss []string = reverseSs(ss)
	var is []int

	for _, v := range rss {
		z, _ := strconv.Atoi(v)
		is = append(is, z)
	}

	for i, v := range is {
		switch i {
		case 0:
			ret = rOne[v]
		case 1:
			ret = rTen[v] + ret
		case 2:
			ret = rHun[v] + ret
		case 3:
			ret = rTho[v] + ret
		}
	}

	return ret
}
