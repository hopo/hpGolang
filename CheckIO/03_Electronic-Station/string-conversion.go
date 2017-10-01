package main

import (
	"fmt"
)

func stepsToConvert(line1, line2 string) (ret int) {
	return
}

func main() {
	fmt.Println(stepsToConvert("line1", "line1")) //  0, "eq"
	fmt.Println(stepsToConvert("line1", "line2")) //  1, "2"
	fmt.Println(stepsToConvert("line", "line2"))  //  1, "none to 2"
	fmt.Println(stepsToConvert("ine", "line2"))   //  2, "need two more"
	fmt.Println(stepsToConvert("line1", "1enil")) //  4, "everything is opposite"
	fmt.Println(stepsToConvert("", ""))           //  0, "two empty"
	fmt.Println(stepsToConvert("l", ""))          //  1, "one side"
	fmt.Println(stepsToConvert("", "l"))          //  1, "another side"
}
