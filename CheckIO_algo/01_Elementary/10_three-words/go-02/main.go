// is it best?

package main

import (
	"fmt"
	"regexp"
)

func threeWords(data string) bool {
	return regexp.MustCompile("[a-zA-Z]+ [a-zA-Z]+ [a-zA-Z]+").MatchString(data)
}

func main() {
	fmt.Println(threeWords("Hello World hello")) //true
	fmt.Println(threeWords("He is 123 man"))     //false
	fmt.Println(threeWords("1 2 3 4"))           //false
	fmt.Println(threeWords("bla bla bla bla"))   //true
	fmt.Println(threeWords("Hi"))                //false
}
