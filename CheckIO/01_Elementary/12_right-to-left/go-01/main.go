package main

import (
	"fmt"
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

	ret := change_word(target, find, word)
	return ret
}

func change_word(t string, f string, w string) string {
	target := []byte(t)
	find := []byte(f)
	word := []byte(w)
	bi := match_str_beginindex(target, find)

	if bi == -1 {
		return string(target)
	}

	flen := len(find)
	ei := bi + flen

	target = append(target[:bi], append([]byte(word), target[ei:]...)...)
	return change_word(string(target), f, w)
}

func match_str_beginindex(tar, chk []byte) int {
	bidx := -1
	var c int

	for i := 0; i < len(tar); i++ {
		if tar[i] == chk[0] {
			for j := 0; ; j++ {
				if tar[i+j] == chk[j] {
					c++
					if c == len(chk) {
						bidx = i
						return bidx
					}
				} else if tar[i+j] != chk[j] {
					c = 0
					break
				}
			}
		}
	}

	return bidx
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
