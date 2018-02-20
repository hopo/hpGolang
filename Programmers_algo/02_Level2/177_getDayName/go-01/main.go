package main

import (
	"fmt"
)

func main() {
	ex1 := getDayName(5, 24) // TUE
	ex2 := getDayName(11, 5) // SAT
	fmt.Println(ex1)
	fmt.Println(ex2)
}

func getDayName(m, d int) string {
	mdays := []int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	dnames := []string{"SUN", "MON", "TUE", "WED", "THU", "FRI", "SAT"}
	for i := 0; i < m-1; i++ {
		d += mdays[i]
	}
	r := d % 7
	var ret string
	if r < 3 {
		ret = dnames[r+4]
	} else {
		ret = dnames[r-3]
	}
	return ret
}
