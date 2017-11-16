package main

import (
	"fmt"
)

func main() {
	ex1 := jaden_case("3people unFollowed me for the last week") // "3people Unfollowed Me For The Last Week"
	ex2 := jaden_case("viDeo kIl1ed radio 5taR")                 // "Video Kil1ed Radio 5tar"
	fmt.Println(ex1)
	fmt.Println(ex2)
}

func jaden_case(s string) string {
	bslsl := split_string(s)
	var ret string
	for _, v := range bslsl {
		for i, _ := range v {
			if isUpper(v[i]) {
				v[i] = v[i] + 32
			}
		}
		if isLower(v[0]) {
			v[0] = v[0] - 32
		}
		ret += string(v)
	}
	return ret
}

func split_string(s string) [][]byte {
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

func isLower(b byte) bool {
	// byte 97~122 is lower case
	for i := 97; i < 123; i++ {
		if int(b) == i {
			return true
		}
	}
	return false
}

func isUpper(b byte) bool {
	// byte 65~90 is upper case
	for i := 65; i < 91; i++ {
		if int(b) == i {
			return true
		}
	}
	return false
}
