package main

import (
	"fmt"
)

func brackets(expression string) interface{} {
	return true || false
}

func main() {
	fmt.Println(brackets("((5+3)*2+1)"))              //true, "Simple"
	fmt.Println(brackets("{[(3+1)+2]+}"))             //true, "Different types"
	fmt.Println(brackets("(3+{1-1)}"))                //false, ") is alone inside {}"
	fmt.Println(brackets("[1+1]+(2*2)-{3/3}"))        //true, "Different operators"
	fmt.Println(brackets("(({[(((1)-2)+3)-3]/3}-3)")) //false, "One is redundant"
	fmt.Println(brackets("2+3"))                      //true, "No brackets, no problem"
}
