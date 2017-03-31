package main

import "strconv"

func fizzBuzz(inpt int) {
	var rslt string
	if inpt%(3*5) == 0 {
		rslt = "Fizz Buzz"
	} else if inpt%3 == 0 {
		rslt = "Fizz"
	} else if inpt%5 == 0 {
		rslt = "Buzz"
	} else {
		t := strconv.Itoa(inpt)
		rslt = t
	}
	println(rslt)
}

func main() {
	fizzBuzz(15) // "Fizz Buzz", "15 is divisible by 3 and 5");
	fizzBuzz(6) // "Fizz", "6 is divisible by 3");
	fizzBuzz(5) // "Buzz", "5 is divisible by 5");
	fizzBuzz(7) // "7", "7 is not divisible by 3 or 5");
}
