package main

import (
	"fmt"
)

func fizzBuzz(n int) string {
	switch {
	case n%15 == 0:
		return fmt.Sprint("Fizz Buzz")
	case n%3 == 0:
		return fmt.Sprint("Fizz")
	case n%5 == 0:
		return fmt.Sprint("Buzz")
	default:
		return fmt.Sprint(n)
	}
}

func main() {
	fmt.Println(fizzBuzz(15)) // "Fizz Buzz", "15 is divisible by 3 and 5");
	fmt.Println(fizzBuzz(6))  // "Fizz", "6 is divisible by 3");
	fmt.Println(fizzBuzz(5))  // "Buzz", "5 is divisible by 5");
	fmt.Println(fizzBuzz(7))  // "7", "7 is not divisible by 3 or 5");
}
