package main

import (
	"fmt"
    "sort"
	"strings"
    "regexp"
)

// func checker(ss []string) string{}

func mostWanted(data string) ([]string, int) {
    // sli := strings.Split(data, "")
    data = strings.ToLower(data)
    sli := regexp.MustCompile(`(?i)[a-z]`).FindAllString(data, -1)
    sort.Strings(sli)
    // r := checker(sli)
    return sli, len(sli)
}

func main() {
    fmt.Println(mostWanted("Hello World!"))	//"l", "1st example"
    fmt.Println(mostWanted("How do you do?"))	//"o", "2nd example"
    fmt.Println(mostWanted("One"))	//"e", "3rd example"
    fmt.Println(mostWanted("Oops!"))	//"o", "4th example"
    fmt.Println(mostWanted("AAaooo!!!!"))	//"a", "Letters"
}
