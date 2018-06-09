// ********
// hpkg hgo
// ********

package hpkg

import (
	"fmt"
	"math"
)

// find and change word in string
func Change_word(t string, f string, w string) string {
	target := []byte(t)
	find := []byte(f)
	word := []byte(w)
	bi := Match_str_beginindex(target, find) // if not found match -> bi = -1

	if bi == -1 {
		return string(target)
	}

	flen := len(find)
	ei := bi + flen

	target = append(target[:bi], append([]byte(word), target[ei:]...)...) // change word
	return Change_word(string(target), f, w)
}

// match between target bsl and check bsl. begin index
func Match_str_beginindex(tar, chk []byte) int {
	bidx := -1
	var c int

	for i := 0; i < len(tar); i++ {
		if tar[i] == chk[0] {
			for j := 0; ; j++ {
				if tar[i+j] == chk[j] { // match check
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

// number base, 16, 10, 5, 2 ... convert
func Number_base(s string, n int) int {
	bsl := []byte(s)
	var isl []int
	for _, v := range bsl {
		if 64 < v && v < 91 {
			isl = append(isl, int(v-55))
		}
		if 47 < v && v < 59 {
			isl = append(isl, int(v-48))
		}
	}

	var j int
	var ret int
	for i, _ := range isl {
		j = len(isl) - 1 - i
		mp := math.Pow(float64(n), float64(i))
		ret += isl[j] * int(mp)
	}
	return ret
}

// slice variable variety maker
func Make_vari_sli(ssl []string) []string {
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
		vri := Make_vari_sli(send)

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
func Remove_samevalue(ssl []string) []string {
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

// input int n, convert (type string)binary
func Binary_convert(n int) string {
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

// string multiple: "*" * 3 = "***"
func Str_multi(s string, n int) (rslt string) {
	for i := 0; i < n; i++ {
		rslt += fmt.Sprint(s)
	}
	return
}

// delete value in int slice by index
func Del_value_isl(isl []int, index int) (ret []int) {
	return append(isl[:index], isl[index+1:]...)
}

// string split to byte slice include blank(byte 32)
func Split_string(s string) [][]byte {
	var ret [][]byte
	var box []byte
	bsl := []byte(s)
	for _, v := range bsl {
		box = append(box, v)
		// include ****32
		if v == 32 {
			ret = append(ret, box)
			box = []byte{}
		}
		if v == bsl[len(bsl)-1] {
			ret = append(ret, box)
			break
		}
	}
	return ret
}

// int split each number
func Split_int(num int) []int {
	numlen := 10
	for {
		if num-numlen < 0 {
			numlen /= 10
			break
		}
		numlen *= 10
	}
	var ret []int
	var ea int
	for {
		ea = num / numlen
		ret = append(ret, ea)
		num %= numlen
		numlen /= 10
		if numlen < 1 {
			break
		}
	}
	return ret
}

// each int in isl make Number
func Isl_makeNum(isl []int) int {
	l := len(isl)
	numlen := Pow(10, l-1) // math power function need
	var ret int
	for _, v := range isl {
		ret += v * numlen
		numlen /= 10
	}
	return ret
}

// isl sort
func Sort_isl(isl []int) []int {
	for i, _ := range isl {
		for j, _ := range isl {
			if i != j && isl[i] < isl[j] { // reverse sort. change '>'
				isl[i], isl[j] = isl[j], isl[i]
			}
		}
	}
	return isl
}

// sort islsl by 0 index
func Sort_islsl(islsl [][]int) [][]int {
	ret := islsl
	for i, v := range ret {
		for j, w := range ret {
			if i != j && v[0] < w[0] {
				ret[i], ret[j] = ret[j], ret[i]
			}
		}
	}
	return ret
}

// member(isl) in islsl true or false
func Member_in_islsl(s [][]int, m []int) bool {
	var chk int
	for _, v := range s {
		for i, w := range v {
			if w != m[i] {
				chk = 0
				break
			} else if w == m[i] {
				chk++
				if chk == len(m) {
					return true // in member(true)
				}
			}
		}
	}
	return false
}
