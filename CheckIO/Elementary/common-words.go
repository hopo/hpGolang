package main

import (
	"fmt"
    "sort"
	"strings"
)

func commonWords(first, second string) {
	frsl := strings.Split(first, ",")
	srsl := strings.Split(second, ",")
    var inter []string

	for _, f := range frsl {
		for _, s := range srsl {
            if f == s {
                inter = append(inter, f)
            }
		}
	}
    sort.Strings(inter)
    rslt := strings.Join(inter, ",")
    fmt.Println(rslt)
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
