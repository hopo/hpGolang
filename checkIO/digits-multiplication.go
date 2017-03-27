package main

import (
	"strconv"
	"strings"
)

func digitsMultip(data int) {
	str := strconv.Itoa(data)
	z := 1
	for i := 0; i < len(str); i++ {
		if str[i] != 0 {
			z *= str[i]
		}
	}
}

func main() {
	digitsMultip(123405) //120
	digitsMultip(999) //729
	digitsMultip(1000) //1
	digitsMultip(1111) //1
}

/*
function digitsMultip(data) {
	var arr = Array.from(String(data));
	var z = 1;
	for(var i=0; i < arr.length; i++){
		if(arr[i] == 0){continue;};
		z = z*arr[i];
	};
    return z;
};
*/
