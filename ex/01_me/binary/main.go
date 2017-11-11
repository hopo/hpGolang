package main

import (
	"fmt"
)

func main() {
	binary(60, 120)
}

// differ between for and for
func binary(a, b int) {
	var ssl []string
	for i := a; i < b; i++ {
		s := fmt.Sprintf("%08b", i)
		ssl = append(ssl, s)
	}
	for i, v := range ssl {
		if i%16 == 0 {
			fmt.Println("")
		}
		r := bPrnt(v)
		fmt.Println(r, ":", i, " : ", btosConv(i))
	}

}

func bPrnt(s string) string {
	if len(s) != 8 {
		return "*** error : input string  is not 8 length"
	}
	bss := []byte(s)
	var a, b string
	for i, v := range bss {
		switch {
		case i < 4:
			a += string(v)
		case i > 3:
			b += string(v)
		}
	}
	return a + " " + b
}

func btosConv(n int) string {
	return string(n)
}
