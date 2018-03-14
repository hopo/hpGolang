// other solution?

package main

import "fmt"

func countInversion(sli []int) int {
	var check int
	for i := 0; i < len(sli); i++ {
		for j := 1; j < len(sli)-i; j++ {
			if sli[i] > sli[i+j] {
				check++
			}
		}
	}
	return check
}

func main() {
	fmt.Println(countInversion([]int{1, 2, 5, 3, 4, 7, 6})) //3
	fmt.Println(countInversion([]int{0, 1, 2, 3}))          //0
	fmt.Println(countInversion([]int{99, -99}))             //1
	fmt.Println(countInversion([]int{5, 3, 2, 1, 0}))       //10
}
