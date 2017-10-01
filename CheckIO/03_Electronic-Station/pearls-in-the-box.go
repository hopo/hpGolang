package main

import (
	"fmt"
)

func boxProbability(marbles string, step float64) (ret float64) {
	return
}

func main() {
	fmt.Println(boxProbability("bbw", 3))     //0.48, "First"
	fmt.Println(boxProbability("wwb", 3))     //0.52, "Second"
	fmt.Println(boxProbability("www", 3))     //0.56, "Third"
	fmt.Println(boxProbability("bbbb", 1))    //0, "Fifth"
	fmt.Println(boxProbability("wwbb", 4))    //0.5, "Sixth"
	fmt.Println(boxProbability("bwbwbwb", 5)) //0.48, "Seventh"
}
