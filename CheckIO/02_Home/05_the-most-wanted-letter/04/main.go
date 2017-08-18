package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func lowSorter(s string) []string {
	low := strings.ToLower(s)
	ss := regexp.MustCompile(`[a-z]`).FindAllString(low, -1)
	sort.Strings(ss)
	return ss
}

func findIndex(ss []string, s string) int {
	len := len(ss)
	for i := 0; i < len; i++ {
		if ss[i] == s {
			return i
		}
	}
	return -1
}

func findLastIndex(ss []string, s string) int {
	len := len(ss)
	for i := len - 1; i > -1; i-- {
		if ss[i] == s {
			return i
		}
	}
	return -1
}

func mostWanted(s string) string {
	ss := lowSorter(s)
	rslt := ss[0]
	n := 0
	for _, v := range ss {
		ea := findLastIndex(ss, v) - findIndex(ss, v)
		if ea > n {
			rslt = v
			n = ea
		}
	}
	return rslt
}

func main() {
	var m1 = "Hello World!"
	var m2 = "How do you do?"
	var m3 = "One"
	var m4 = "Oops!"
	var m5 = "AAaooo!!!!"

	fmt.Println(mostWanted(m1)) //"l", "1st example"
	fmt.Println(mostWanted(m2)) //"o", "2nd example"
	fmt.Println(mostWanted(m3)) //"e", "3rd example"
	fmt.Println(mostWanted(m4)) //"o", "4th example"
	fmt.Println(mostWanted(m5)) //"a", "Letters"
}
