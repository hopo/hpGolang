package main

import (
	"fmt"
	"reflect"
	// "regexp"
	// "strconv"
	// "strings"
)

func main() {
	data := "abcde"
	t := reflect.TypeOf(data)
	fmt.Println(t)
}

/*
func regExpression(){
	data := "He is 123 man"

	re := regexp.MustCompile("[a-zA-Z]+")
	rslt := re.FindAllString(data, -1) //-1:all
	fmt.Println(rslt)
	fmt.Printf("rslt: %T\n", rslt)

	mat := re.MatchString(data)
	fmt.Println(mat)
}
*/

/*
//convert between str and num
func strConvert() {
	num := 1234
	x := strconv.Itoa(num)

	i, _ := strconv.Atoi(x)

	println(x, i)
	fmt.Printf("x: %T, i: %T", x, i)
}
*/
