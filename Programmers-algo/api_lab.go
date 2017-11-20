// *** api lab

package main

import (
	"fmt"
)

func main() {
	//fmt.Println(binary_convert(78)) // 1001110
	make_vari([]string{"0", "1", "2"})
}

// variety
func make_vari(ssl []string) {
	vari3 := variety3(ssl)
	fmt.Println(vari3)
}

func variety3(ssl []string) []string {
	var fix string
	var ret, box []string
	box = append(box, ssl...)
	lnth := len(box)
	for i := 0; i < lnth; i++ {
		fix = box[i]
		box[i] = ""
		var send []string
		for _, v := range box {
			if v != "" {
				send = append(send, v)
			}
		}

		vr2 := variety2(send)
		var s string
		var rcv []string
		for _, v := range vr2 {
			s += v
			rcv = append(rcv, s)
			s = ""
		}

		for _, v := range rcv {
			s = fix + v
			ret = append(ret, s)
			s = ""
		}
		box[i] = fix
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

// input int n, convert (type string)binary
func binary_convert(n int) string {
	//fmt.Println("n int:", n) // checker
	//fmt.Printf("%b\n", n) // checker
	var tsl []int
	for i := 1; i < n+1; i *= 2 {
		tsl = append([]int{i}, tsl...) // push front
	}

	zno := make([]int, len(tsl))
	for i := 0; i < len(tsl); i++ {
		n = n - tsl[i]
		if n == 0 {
			zno[i] = 1
			break
		}
		if n > 0 {
			zno[i] = 1
		}
		if n < 0 {
			n = n + tsl[i]
		}
	}

	var r string
	for _, v := range zno {
		r += string(v + 48) // int make string
	}
	return r
}
