// ...ing
package main

import (
	"fmt"
)

func main() {
	//ex1 := between_markers("What is >apple<", ">", "<") // "apple", "One sym"
	//fmt.Println(ex1)

	ex2 := between_markers("<head><title>My new site</title></head>", "<title>", "</title>") // "My new site", "HTML"
	fmt.Println(ex2)
	/*
		ex3 := between_markers("No[/b] hi", "[b]", "[/b]") // "No", 'No opened'
		fmt.Println(ex3)
		ex4 := between_markers("No [b]hi", "[b]", "[/b]") // "hi", 'No close'
		fmt.Println(ex4)
		ex5 := between_markers("No hi", "[b]", "[/b]") // "No hi", 'No markers at all'
		fmt.Println(ex5)
		ex6 := between_markers("No <hi>", ">", "<") // "", 'Wrong direction'
		fmt.Println(ex6)
	*/
}

func between_markers(text, begin, end string) string {
	btext := []byte(text)
	bbeg := []byte(begin)
	//bend := []byte(end)
	for i, _ := range btext {
		for j, _ := range bbeg {
			if btext[i+j] == bbeg[j] {
				fmt.Println(i + j)
			}
		}
	}
	return text
}
