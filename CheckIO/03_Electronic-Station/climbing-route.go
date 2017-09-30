package main

import "fmt"

func climbingRoute(elevation_map []string) (ret int) {
	return
}

func main() {
	fmt.Println(climbingRoute([]string{
		"000",
		"210",
		"000"})) // 6, 'basic'
	fmt.Println(climbingRoute([]string{
		"00000",
		"05670",
		"04980",
		"03210",
		"00000"})) // 26, 'spiral'
	fmt.Println(climbingRoute([]string{
		"000000001",
		"222322222",
		"100000000"})) // 26, 'bridge'
	fmt.Println(climbingRoute([]string{
		"000000002110",
		"011100002310",
		"012100002220",
		"011100000000"})) // 26, 'two top'
	fmt.Println(climbingRoute([]string{
		"000000120000",
		"001002432100",
		"012111211000",
		"001000000000"})) // 16, 'one top'
	fmt.Println(climbingRoute([]string{
		"00000000111111100",
		"00000000122222100",
		"00000000123332100",
		"00000000123432100",
		"00000000123332100",
		"00000000122222100",
		"00000000111111100",
		"00011111000000000",
		"00012221000000000",
		"00012321000000000",
		"00012221000000012",
		"00011111000000000",
		"11100000000000000",
		"12100000000000000",
		"11100000000000000"})) // 52, 'pyramids'
}
