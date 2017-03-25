package main

func fizzBuzz(inpt int) {
	rslt := inpt
	println(rslt)
}

func main() {
	println("* * * * * * * * * *")
	fizzBuzz(15) // "Fizz Buzz", "15 is divisible by 3 and 5");
	fizzBuzz(6) // "Fizz", "6 is divisible by 3");
	fizzBuzz(5) // "Buzz", "5 is divisible by 5");
	fizzBuzz(7) // "7", "7 is not divisible by 3 or 5");
	println("* * * * * * * * * *")
}
