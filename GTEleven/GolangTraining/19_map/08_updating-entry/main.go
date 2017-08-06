package main

import "fmt"

func main() {

	myGreeting := map[string]string{
		"Tim":   "Good mornig!",
		"Jenny": "Bonjour!",
	}

	myGreeting["Harleen"] = "Howdy"
	fmt.Println(myGreeting)
	myGreeting["Harleen"] = "Gidday"
	fmt.Println(myGreeting)
}
