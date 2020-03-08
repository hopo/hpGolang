// Go's `math/rand` package provides
// [pseudorandom number](http://en.wikipedia.org/wiki/Pseudorandom_number_generator)
// generation.

package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	for {
		num := rand.Intn(100)
		fmt.Println()
		fmt.Print(num*num, " ???  ")

		var a string
		fmt.Scanln(&a)

		v, _ := strconv.Atoi(a)

		switch {
		case a == "q":
			return
		case num == v:
			fmt.Println("Good Job!")
		case num != v:
			fmt.Println("no: answer is ", num)
		}
	}
}
