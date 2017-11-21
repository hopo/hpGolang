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

	dd := dduim(n) // dduim make, [2 2 1]
	fmt.Println("dd:", dd)
	sdd := [][]int{dd}         // data for two_breaker()
	vrdd_i := two_breaker(sdd) // dd two breaker variety dduim slice
	fmt.Println("vrdd_i", vrdd_i)

	vrdd_s := isl_conv_ssl(vrdd_i) // data conv for make_vari_sli()
	fmt.Printf("vrdd_s: %T, %v\n", vrdd_s, vrdd_s)

	vri := make_vari_sli(vrdd_s[0])
	fmt.Println("vri:", vri)

	rmvd := remove_samevalue(vri)
	fmt.Println("rmvd:", rmvd)

	return -1
}

func dduim(n int) []int {
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

func two_breaker(slsl [][]int) [][]int {
	ret := make([][]int, len(slsl))
	copy(ret, slsl)
	chk := make([]int, len(ret[len(ret)-1]))
	copy(chk, ret[len(ret)-1])
	var j int
	for i := 0; i < len(chk); i++ {
		j = len(chk) - 1 - i
		if chk[j] == 2 {
			chk[j] = 1
			chk = append(chk, 1)
			ret = append(ret, chk)
			return two_breaker(ret)
		}
	}
	return ret
}

func isl_conv_ssl(islsl [][]int) [][]string {
	var ret [][]string
	var ssl []string
	for _, dd := range islsl {
		for _, v := range dd {
			ssl = append(ssl, string(v+48))
		}
		ret = append(ret, ssl)
		ssl = []string{}
	}
	return ret
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
