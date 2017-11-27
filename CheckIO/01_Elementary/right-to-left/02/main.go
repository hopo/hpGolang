package main

import (
	"fmt"
	"github.com/hopo/hpkg"
)

func main() {
	ex1 := left_join([]string{"left", "right", "left", "stop"}) // "left,left,left,stop", "All to left"
	fmt.Println(ex1)
	ex2 := left_join([]string{"bright aright", "ok"}) // "bleft aleft,ok", "Bright Left"
	fmt.Println(ex2)
	ex3 := left_join([]string{"brightness wright", ""}) // "bleftness wleft", "One phrase"
	fmt.Println(ex3)
	ex4 := left_join([]string{"enough", "jokes"}) // "enough,jokes", "Nothing to replace"
	fmt.Println(ex4)
}

func left_join(phr []string) string {
	target := ssl_join_str(phr)
	find := "right" // want to find string "right"
	word := "left"

	ret := hpkg.Change_word(target, find, word)

	return ret
}

func ssl_join_str(phr []string) string {
	var jnstr string

	for i, v := range phr {
		if v == "" {
			x := []byte(jnstr)
			x = x[:len(x)-1]
			jnstr = string(x)
			break
		}
		if i == len(phr)-1 {
			jnstr += v
			break
		}
		jnstr += v + ","
	}
	return jnstr
}
