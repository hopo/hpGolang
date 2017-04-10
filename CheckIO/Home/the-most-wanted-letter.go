package main

import (
	"fmt"
    "sort"
	"strings"
    "regexp"
)

func str2sli(str string) []string {
    str = strings.ToLower(str)
    r := regexp.MustCompile(`(?i)[a-z]`).FindAllString(str, -1)
    sort.Strings(r)
    return r
}

func checker(sli []string) string {
    var str string
    for _, s := range sli { str += s }

    r := sli[0]
    z := 1
    for _, s := range sli {
        n := strings.Count(str, s) //version1
        // n := strings.LastIndex(str, s)-strings.Index(str, s) //verion2
        if z < n {
            r = s
            z = n
        }
    }
    return r
}

func mostWanted(data string) string {
    return checker(str2sli(data))
}

func main() {
    fmt.Println(mostWanted("Hello World!"))	//"l", "1st example"
    fmt.Println(mostWanted("How do you do?"))	//"o", "2nd example"
    fmt.Println(mostWanted("One"))	//"e", "3rd example"
    fmt.Println(mostWanted("Oops!"))	//"o", "4th example"
    fmt.Println(mostWanted("AAaooo!!!!"))	//"a", "Letters"
}
