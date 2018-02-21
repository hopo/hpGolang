package main

import (
	"fmt"
)

func main() {
	ex1 := between_markers("What is >apple<", ">", "<") // "apple", "One sym"
	fmt.Println(ex1)
	ex2 := between_markers("<head><title>My new site</title></head>", "<title>", "</title>") // "My new site", "HTML"
	fmt.Println(ex2)
	ex3 := between_markers("No[/b] hi", "[b]", "[/b]") // "No", 'No opened'
	fmt.Println(ex3)
	ex4 := between_markers("No [b]hi", "[b]", "[/b]") // "hi", 'No close'
	fmt.Println(ex4)
	ex5 := between_markers("No hi", "[b]", "[/b]") // "No hi", 'No markers at all'
	fmt.Println(ex5)
	ex6 := between_markers("No <hi>", ">", "<") // "", 'Wrong direction'
	fmt.Println(ex6)
}

func between_markers(text, begin, end string) string {
	var ret []byte
	bsl := []byte(text)
	oi := find_oi(text, begin)
	ci := find_ci(text, end)

	if oi < 0 && ci < 0 {
		ret = bsl
	} else if ci < 0 {
		ret = bsl[oi+1:]
	} else if oi > ci {
		ret = nil
	} else {
		ret = bsl[oi+1 : ci]
	}
	return string(ret)
}

func find_oi(text string, begin string) int {
	var x, y byte
	var chk int
	oi := -1
	for i, _ := range text {
		if text[i] == begin[0] {
			for j, _ := range begin {
				x, y = text[i+j], begin[j]
				if x == y {
					chk++
					if chk == len(begin) {
						oi = i + j
						return oi
					}
				} else {
					chk = 0
					break
				}
			}
		}
	}
	return oi
}

func find_ci(text string, end string) int {
	var x, y byte
	var chk int
	ci := -1
	for i, _ := range text {
		if text[i] == end[0] {
			for j, _ := range end {
				x, y = text[i+j], end[j]
				if x == y {
					chk++
					if chk == len(end) {
						ci = i
						return ci
					}
				}
			}
		}
	}
	return ci
}
