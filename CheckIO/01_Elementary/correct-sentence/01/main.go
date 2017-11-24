package main

import (
	"fmt"
)

func main() {
	ex1 := correct_sentence("greeting, friends") // "Greeting, friends."
	fmt.Println(ex1)
	ex2 := correct_sentence("Greeting, friends") // "Greeting, friends."
	fmt.Println(ex2)
	ex3 := correct_sentence("Greeting, friends.") // "Greeting, friends."
	fmt.Println(ex3)
	ex4 := correct_sentence("hi") // "Hi."
	fmt.Println(ex4)
}

func correct_sentence(text string) string {
	bsl := []byte(text)
	if 96 < bsl[0] && bsl[0] < 123 {
		bsl[0] = text[0] - 32
	}
	if bsl[len(bsl)-1] != 46 {
		bsl = append(bsl, byte(46))
	}
	return string(bsl)
}
