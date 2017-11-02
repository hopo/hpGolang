package main

import "fmt"

func main() {
	arr := []string{"hello"}
	fmt.Println(arr)

	arr = append(arr, "go-go-go-")
	fmt.Println(arr)
}
