package main

import (
	"fmt"
	"strconv"
)

func main() {
	ex1 := say_hi("Alex", 32) // "Hi. My name is Alex and I'm 32 years old", "First"
	fmt.Println(ex1)
	ex2 := say_hi("Frank", 68) // "Hi. My name is Frank and I'm 68 years old", "Second"
	fmt.Println(ex2)
}

func say_hi(name string, age int) string {
	agestr := strconv.Itoa(age)
	return "Hi. My name is " + name + " and I'm " + agestr + " years old"
}
