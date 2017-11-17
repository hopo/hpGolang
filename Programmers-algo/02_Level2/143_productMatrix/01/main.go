package main

import (
	"fmt"
)

func main() {
	ex1 := productMatrix([][]int{{1, 2}, {2, 3}}, [][]int{{3, 4}, {5, 6}})
	fmt.Println(ex1)

}

func productMatrix(A, B [][]int) [][]int {
	if len(A[0]) != len(B) {
		fmt.Println(" *** Operating impossible")
		return nil
	}

	l := len(A[0])
	var x, y int
	var xsl, ysl []int
	var xslsl, yslsl [][]int
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			x = A[i][0] * B[0][j]
			y = A[i][1] * B[1][j]
			xsl = append(xsl, x)
			ysl = append(ysl, y)
		}
		xslsl = append(xslsl, xsl)
		yslsl = append(yslsl, ysl)
		xsl, ysl = []int{}, []int{}
	}

	var sumsl []int
	var ret [][]int
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			sum := xslsl[i][j] + yslsl[i][j]
			sumsl = append(sumsl, sum)
		}
		ret = append(ret, sumsl)
		sumsl = []int{}
	}

	return ret
}
