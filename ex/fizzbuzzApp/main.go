package main

import "fmt"

func fizzbuzz(i int) {
	if i%15 == 0 {
		fmt.Println(i, " -- FizzBuzz")
	} else if i%3 == 0 {
		fmt.Println(i, " -- FIZZ")
	} else if i%5 == 0 {
		fmt.Println(i, " -- BUZZ")
	} else {
		fmt.Println(i, "?")
	}
}

func main() {
	var k int
	fmt.Println("FizzBuzz? ")
	fmt.Scan(&k)
	fizzbuzz(k)
}
