package main

import (
	"fmt"
    // "sort"
	"strings"
    "regexp"
)

func mostWanted(data string) []string {
    // sli := strings.Split(data, "")

    str := regexp.MustCompile(`(?i)[^a-z]`).ReplaceAllString(data, "")
    str = strings.ToLower(str)
    a := regexp.MustCompile("[o]+").FindAllString(str, -1)
    return a
}

func main() {
    fmt.Println(mostWanted("Hello World!"))	//"l", "1st example"
    fmt.Println(mostWanted("How do you do?"))	//"o", "2nd example"
    fmt.Println(mostWanted("One"))	//"e", "3rd example"
    fmt.Println(mostWanted("Oops!"))	//"o", "4th example"
    fmt.Println(mostWanted("AAaooo!!!!"))	//"a", "Letters"
}
