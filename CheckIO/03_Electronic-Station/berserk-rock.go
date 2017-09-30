package main

import (
	"fmt"
)

func berserkRook(marbles string, step []string) interface{} {
	x := append(step, marbles)
	return x
}

func main() {
	fmt.Println(berserkRook("d3", []string{"d6", "b6", "c8", "g4", "b8", "g6"})) //5, "First"
	fmt.Println(berserkRook("a2", []string{"f6", "f2", "a6", "f8", "h8", "h6"})) //6, "Second"
	fmt.Println(berserkRook("a2", []string{"f6", "f8", "f2", "a6", "h6"}))       //4, "Third"
}
