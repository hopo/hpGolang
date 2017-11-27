// is it best?

package main

import (
	"fmt"
	"regexp"
)

func findMessage(data string) string {
	return regexp.MustCompile("[^A-Z]").ReplaceAllString(data, "")
}

func main() {
	fmt.Println(findMessage("How are you? Eh, ok. Low or Lower? Ohhh.")) //"HELLO", "hello"
	fmt.Println(findMessage("hello world!"))                             //"", "Nothing"
	fmt.Println(findMessage("HELLO WORLD!!!"))                           //"HELLOWORLD", "Capitals"
}
