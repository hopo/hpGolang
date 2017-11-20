package main

import (
	"fmt"
)

func main() {
	ex1 := jumpCase(5) // 5
	fmt.Println(ex1)
}

func jumpCase(n int) int {
	fmt.Println("n:", n)

	isl := big_dduim(n)
	var ssl []string
	for _, v := range isl {
		ssl = append(ssl, string(v+48))
	}
	vri := make_vari_sli(ssl)
	fmt.Println("vri:", vri)

	rmvd := remove_samething(vri)
	var ret []string
	ret = append(ret, rmvd...)

	return len(ret)
}

func big_dduim(n int) []int {
	var isl []int
	for {
		if n == 2 {
			isl = append(isl, n)
			break
		}
		if n == 1 {
			isl = append(isl, n)
			break
		}
		if n-2 > 0 {
			isl = append(isl, 2)
			n -= 2
		}
	}
	return isl
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

func remove_samething(ssl []string) []string {
	var ret []string

	fmt.Println("rmvd ret:", ret)

	return ret
}
