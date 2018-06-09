package main

import (
	"fmt"
)

func main() {
	sl := []string{"a", "b", "c"}
	x := "x"

	pop(sl)
	pop_back(sl)
	push_front(sl, x)
	shift(sl)
	unshift(sl, x)
}

func pop(sl []string) {
	var x string
	x, sl = sl[0], sl[1:]
	fmt.Print("pop: ")
	fmt.Println(x, sl)
}

func pop_back(sl []string) {
	var x string
	lnth := len(sl)
	x, sl = sl[lnth-1], sl[:lnth-1]
	fmt.Print("pop back: ")
	fmt.Println(x, sl)
}

func push_front(sl []string, x string) {
	sl = append([]string{x}, sl...)
	fmt.Print("push_front: ")
	fmt.Println(x, sl)

}

func shift(sl []string) {
	x, sl := sl[0], sl[1:]
	fmt.Print("shift: ")
	fmt.Println(x, sl)
}

func unshift(sl []string, x string) {
	sl = append([]string{x}, sl...)
	fmt.Print("unshift: ")
	fmt.Println(x, sl)
}
