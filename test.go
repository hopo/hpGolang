package main

import "time"

func main() {
	t1 := time.Now().Weekday()
	t2 := time.Now().Hour()
	println(t1, t2)
}

