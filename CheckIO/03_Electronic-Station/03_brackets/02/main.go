//anothet solution?
package main

import (
	"fmt"
	"regexp"
)

func brackets(expression string) (ret bool) {
	re := regexp.MustCompile(`[^0-9\+\-\*\/]`) // `/D` : not digit
	ss := re.FindAllString(expression, -1)

	fmt.Print("ss: ", ss, " @")

	l := len(ss)
	if l == 0 {
		ret = true
		return
	} else if l%2 != 0 {
		fmt.Print(" *Not Fair* ")
		return
	}

	f := ss[0]
	e := ss[l-1]

	bra := map[string]string{"(": ")", "[": "]", "{": "}"}
	for k, v := range bra {
		if f == v {
			fmt.Print(" *Not Open* ")
			return
		}
		if e == k {
			fmt.Print(" *Not Close* ")
			return
		}
	}

	var is []int
	for i := 1; i < l; i++ {
		if i%2 == 1 && bra[f] == ss[i] {
			is = append(is, i)
		}
	}
	if is == nil {
		fmt.Print(" *Alone Inside* ")
		return
	}

	ret = true
	return
}

func main() {
	fmt.Println(brackets("(2+3)-{3*2{"))              //false, "hp ex"
	fmt.Println(brackets("((5+3)*2+1)"))              //true, "Simple"
	fmt.Println(brackets("{[(3+1)+2]+}"))             //true, "Different types"
	fmt.Println(brackets("(3+{1-1)}"))                //false, ") is alone inside {}"
	fmt.Println(brackets("[1+1]+(2*2)-{3/3}"))        //true, "Different operators"
	fmt.Println(brackets("(({[(((1)-2)+3)-3]/3}-3)")) //false, "One is redundant"
	fmt.Println(brackets("2+3"))                      //true, "No brackets, no problem"
}
