package main

import (
	"fmt"
	"strings"
)

func commonWords(first, second string) {
	frsl := strings.Split(first, ",")
	srsl := strings.Split(second, ",")

	for _, f := range frsl {
		for _, s := range srsl {
			fmt.Println(f, s)
		}
	}
}

func main() {
    commonWords("hello,world", "hello,earth")	//"hello"
	commonWords("one,two,three", "four,five,six")	//""
    commonWords("one,two,three", "four,five,one,two,six,three")	//"one,three,two"
}

/*
function commonWords(first, second) {
	var intersect = [];
	for(var f of first.split(",")){
		for(var s of second.split(",")){
			if(f == s){intersect.push(f)};
		};
	};
	return String(intersect.sort());
};
*/
