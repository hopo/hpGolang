package main

import (
// "fmt"
// "math"
// "reflect"
// "regexp"
// "sort"
// "strconv"
// "strings"
)

/*
type ByAbs []float64
func (p ByAbs) Len() int { return len(p) }
func (p ByAbs) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p ByAbs) Less(i, j int) bool { return math.Abs(p[i]) < math.Abs(p[j]) }

func absoluteSorting(nums []float64) {
    sort.Sort(ByAbs(nums))
    fmt.Println(nums)
}
*/

/*
func sortSlice(nums ...float64) {
    fmt.Println(nums)
    fmt.Println(reflect.TypeOf(nums))
    sort.Float64s(nums) // no return value
    fmt.Println(nums)
}
*/

/*
func regExpression(data string){
	re := regexp.MustCompile("[a-zA-Z]+")
	rslt := re.FindAllString(data, -1) //-1:all
	fmt.Println(rslt)
	fmt.Printf("rslt: %T\n", rslt)

	mat := re.MatchString(data)
	fmt.Println(mat)
}
*/

/*
func ConvertStrInt(num int) {
	x := strconv.Itoa(num)
	i, _ := strconv.Atoi(x)

	println(x, i)
	fmt.Printf("x: %T, i: %T", x, i)
}
*/
