package main

import (
    "fmt"
	"strconv"
	"strings"
)

func digitsMultip(data int) int {
	str := strconv.Itoa(data)
	arr := strings.Split(str, "")
	z := 1
	for i := 0; i < len(arr); i++ {
		n, _ := strconv.Atoi(arr[i])
		if n != 0 {
			z *= n
		}
	}
	return z
}

func main() {
	fmt.Println(digitsMultip(123405)) //120
	fmt.Println(digitsMultip(999)) //729
	fmt.Println(digitsMultip(1000)) //1
	fmt.Println(digitsMultip(1111)) //1
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
