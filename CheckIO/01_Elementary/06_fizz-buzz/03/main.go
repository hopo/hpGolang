// to edit

package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) string {
	var res []string
	switch {
	case n%3 == 0:
		res = append(res, "Fizz")
	case n%5 == 0:
		res = append(res, "Buzz")
	default:
		s := strconv.Itoa(n)
		res = append(res, s)
	}
	return fmt.Sprint(res)
}

func main() {
	fmt.Println(fizzBuzz(15)) // "Fizz Buzz", "15 is divisible by 3 and 5");
	fmt.Println(fizzBuzz(6))  // "Fizz", "6 is divisible by 3");
	fmt.Println(fizzBuzz(5))  // "Buzz", "5 is divisible by 5");
	fmt.Println(fizzBuzz(7))  // "7", "7 is not divisible by 3 or 5");
}
