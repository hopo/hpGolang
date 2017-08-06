package main

import "fmt"

func main() {

	var myGreeting = make(map[string]string)
	myGreeting["Tim"] = "Good mornig."
	myGreeting["Jenny"] = "Bonjour."

	fmt.Println(myGreeting)
}
