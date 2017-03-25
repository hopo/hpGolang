package main

import "fmt"

//"""v2
func split(sum int) (int, int) {
	x := sum * 4 / 9
	y := sum - x
	return x, y
}

// //"""v1
// func split(sum int) (x, y int) {
// 	x = sum * 4 / 9
// 	y = sum - x
// 	return
// }

func main() {
	fmt.Println(split(17))
}
