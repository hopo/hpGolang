package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	prnti("s")
}

func prnti2(i interface{}) interface{} {
	switch i.(type) {
	case int:
		fmt.Println("int!")
	case string:
		fmt.Println("string!")
	}
	return 0
}

func prnti(i interface{}) interface{} {
	//t := fmt.Sprintf("%T", i)
	t := reflect.TypeOf(i).Kind()

	fmt.Printf("i: %T ~ %v\n", i, i)
	fmt.Printf("t: %T ~ %v\n", t, t)
	return 0
}

func zeroMulti() {
	z := "0"
	d := 3

	var s string
	for i := 0; i < d; i++ {
		s += z
	}
	fmt.Println(s)

}

func baseConvert() {
	v := int64(10)
	w := int64(9)

	sv := strconv.FormatInt(v, 2)
	fmt.Printf("%T, %v\n", sv, sv)

	sw := strconv.FormatInt(w, 2)
	fmt.Printf("%T, %v\n", sw, sw)

	z1, _ := strconv.ParseFloat(sv, 32)
	z2, _ := strconv.ParseFloat(sw, 32)

	ret := z1 - z2

	fmt.Printf("ret :%T, %v", ret, ret)
}
