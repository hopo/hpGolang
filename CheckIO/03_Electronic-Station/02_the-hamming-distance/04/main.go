package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	fmt.Println(hammingDistance(117, 17))      //3, "First example"
	fmt.Println(hammingDistance(1, 2))         //2, "Second example"
	fmt.Println(hammingDistance(16, 15))       //5, "Third example"
	fmt.Println(hammingDistance("abc", "aaa")) //2, "hp ex""
}

func hammingDistance(a, b interface{}) (ret int) {
	if a == b {
		return 0
	}

	ta := reflect.TypeOf(a).Kind()
	tb := reflect.TypeOf(b).Kind()
	if ta != tb {
		//fmt.Printf("ta: %v, tb: %v\n", ta, tb)
		return -1
	}

	switch a.(type) {
	case int:
		ret = measureInt(a.(int), b.(int))
	case string:
		ret = measureStr(a.(string), b.(string))
	}

	return
}

func measureInt(n1, n2 int) (dist int) {
	x := bConverter(n1)
	y := bConverter(n2)

	lx := len(x)
	ly := len(y)

	var l int

	switch {
	case lx == ly:
		l = lx
	case lx != ly:
		switch {
		case lx > ly:
			d := lx - ly
			y = zMaker(d, "0") + y
			l = lx
		case lx < ly:
			d := ly - lx
			x = zMaker(d, "0") + x
			l = ly
		}
	}

	for i := 0; i < l; i++ {
		if x[i] != y[i] {
			dist++
		}
	}

	return
}

func measureStr(s1, s2 string) (dist int) {
	x := s1
	y := s2

	lx := len(x)
	ly := len(y)

	var l int

	switch {
	case lx == ly:
		l = lx
	case lx != ly:
		switch {
		case lx > ly:
			d := lx - ly
			y = y + zMaker(d, ".")
			l = lx
		case lx < ly:
			d := ly - lx
			x = x + zMaker(d, ".")
			l = ly
		}
	}

	for i := 0; i < l; i++ {
		if x[i] != y[i] {
			dist++
		}
	}

	return
}

func bConverter(num int) (s string) {
	s = strconv.FormatInt(int64(num), 2)
	return
}

func zMaker(m int, z string) (s string) {
	for i := 0; i < m; i++ {
		s += z
	}
	return
}
