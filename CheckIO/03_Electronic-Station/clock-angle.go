package main

import (
	"fmt"
)

func clockAngle(time string) int {
	return 0
}

func main() {
	fmt.Println(clockAngle("02:30")) //105, "02:30"
	fmt.Println(clockAngle("13:42")) //159, "13:42"
	fmt.Println(clockAngle("01:42")) //159, "01:42"
	fmt.Println(clockAngle("01:43")) //153.5, "01:43"
	fmt.Println(clockAngle("00:00")) //0, "Zero"
	fmt.Println(clockAngle("12:01")) //5.5, "Little later"
	fmt.Println(clockAngle("18:00")) //180, "Opposite"
}
