package main

import "fmt"

func brokenClock(startingTime, wrongTime, errorDescription string) interface{} {
	return "00:59:59"
}

func main() {
	fmt.Println(brokenClock("00:00:00", "00:00:15", "+5 seconds at 10 seconds")) // '00:00:10', "First example"
	fmt.Println(brokenClock("06:10:00", "06:10:15", "-5 seconds at 10 seconds")) // '06:10:30', "Second example"
	fmt.Println(brokenClock("13:00:00", "14:01:00", "+1 second at 1 minute"))    // '14:00:00', "Third example"
	fmt.Println(brokenClock("01:05:05", "04:05:05", "-1 hour at 2 hours"))       // '07:05:05', "Fourth example"
	fmt.Println(brokenClock("00:00:00", "00:00:30", "+2 seconds at 6 seconds"))  // '00:00:22', "Fifth example"
}
