// another solution?

package main

import (
	"fmt"
	"sort"
	"strings"
)

func commonWords(first, second string) string {
	frsl := strings.Split(first, ",")
	srsl := strings.Split(second, ",")
	var inter []string

	for _, f := range frsl {
		for _, s := range srsl {
			if f == s {
				inter = append(inter, f)
			}
		}
	}
	sort.Strings(inter)
	rslt := strings.Join(inter, ",")
	return rslt
}

func main() {
	fmt.Println(commonWords("hello,world", "hello,earth"))                   //"hello"
	fmt.Println(commonWords("one,two,three", "four,five,six"))               //""
	fmt.Println(commonWords("one,two,three", "four,five,one,two,six,three")) //"one,three,two"
}
