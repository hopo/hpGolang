// *** api lab

package main

import (
	"fmt"
)

func main() {
	ex := make_vari_sli([]string{"x", "y", "x"})
	fmt.Println(ex)

	ex2 := remove_samevalue(ex)
	fmt.Print(ex2)
}

// slice variable variety maker
func make_vari_sli(ssl []string) []string {
	var fix string
	var ret, box []string
	box = ssl
	lnth := len(ssl)
	if lnth == 1 {
		return ssl
	}

	for i := 0; i < lnth; i++ {
		fix, box = box[0], box[1:] // set fix, box

		var send []string
		send = make([]string, len(box))
		copy(send, box)
		vri := make_vari_sli(send)

		var rcv string
		for _, v := range vri {
			rcv = fix + v
			ret = append(ret, rcv)
			rcv = ""
		}

		box = append(box, fix) // init: fix, box
	}
	return ret
}

// remove same value in ssl
func remove_samevalue(ssl []string) []string {
	var box, ret []string
	box = ssl
	lnth := len(box)
	var tag string
	for i := 0; i < lnth; i++ {
		tag = box[i] // set
		for j := i; j < lnth; j++ {
			if i != j && tag == box[j] {
				box[i] = ""
				break
			}
		}
	}
	for _, v := range box {
		if v != "" {
			ret = append(ret, v)
		}
	}
	return ret
}
