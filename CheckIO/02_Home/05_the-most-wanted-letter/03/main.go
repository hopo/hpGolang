// go new solution

package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

type Data struct {
	rslt string
	str  string
	sli  []string
}

func (d *Data) sorter() {
	str := strings.ToLower(d.str)
	sl := regexp.MustCompile(`(?i)[a-z]`).FindAllString(str, -1)
	sort.Strings(sl)
	d.sli = sl

	var st string
	for _, s := range sl {
		st += s
	}
	d.str = st
}

func (d *Data) checker() {
	r := d.sli[0]
	z := 1
	for _, s := range d.sli {
		n := strings.LastIndex(d.str, s) - strings.Index(d.str, s)
		if z < n {
			r = s
			z = n
		}
	}
	d.rslt = r
}

func mostWanted(data string) string {
	dt := Data{str: data}
	dt.sorter()
	dt.checker()
	return dt.rslt
}

func main() {
	fmt.Println(mostWanted("Hello World!"))   //"l", "1st example"
	fmt.Println(mostWanted("How do you do?")) //"o", "2nd example"
	fmt.Println(mostWanted("One"))            //"e", "3rd example"
	fmt.Println(mostWanted("Oops!"))          //"o", "4th example"
	fmt.Println(mostWanted("AAaooo!!!!"))     //"a", "Letters"
}
