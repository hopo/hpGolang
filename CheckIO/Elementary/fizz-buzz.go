package main

import (
    "fmt"
    "strconv"
)

func fizzBuzz(inpt int) string {
	if inpt%(3*5) == 0 { return "Fizz Buzz"	}
    if inpt%3 == 0 { return "Fizz" }
    if inpt%5 == 0 { return "Buzz" }
	return strconv.Itoa(inpt)
}

func main() {
	fmt.Println(fizzBuzz(15)) // "Fizz Buzz", "15 is divisible by 3 and 5");
	fmt.Println(fizzBuzz(6)) // "Fizz", "6 is divisible by 3");
	fmt.Println(fizzBuzz(5)) // "Buzz", "5 is divisible by 5");
	fmt.Println(fizzBuzz(7)) // "7", "7 is not divisible by 3 or 5");
}
