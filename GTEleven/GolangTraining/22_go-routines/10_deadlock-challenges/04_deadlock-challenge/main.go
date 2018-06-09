package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	for {
		fmt.Println(<-c)
	}
}

// This prints the number, bur we have received this error:
// fatal error: all goroutines are asleep - deadlock!
// where is the deadlock?
// Why are all goroutines asleep?
// How can we dix this?
