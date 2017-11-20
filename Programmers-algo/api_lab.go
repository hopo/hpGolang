// *** api lab

package main

import (
	"fmt"
)

func main() {
	//fmt.Println(binary_convert(78)) // 1001110
	make_vari([]string{"a", "b", "c"})
}

// variety
func make_vari(ssl []string) {
	vari3 := variety3(ssl)
	fmt.Println(vari3)
}

func variety3(ssl []string) []string {
	var fix string
	var ret, box []string
	lnth := len(ssl)
	//	box = make([]string, lnth)
	//	copy(box, ssl)
	box = ssl

	for i := 0; i < lnth; i++ {
		fix, box = box[0], box[1:] // set: fix, box

		var send []string
		send = make([]string, len(box))
		copy(send, box)
		vr2 := variety2(send)

		var rcv string
		for _, v := range vr2 {
			rcv = fix + v
			ret = append(ret, rcv)
			rcv = ""
		}

		box = append(box, fix) // init: fix, box
	}

	return ret
}

func variety2(ssl []string) []string {
	var s string
	var ret []string
	for _, v := range ssl {
		s += v
	}
	ret = append(ret, s)
	s = ""
	ssl[0], ssl[1] = ssl[1], ssl[0]
	for _, v := range ssl {
		s += v
	}
	ret = append(ret, s)
	return ret
}
