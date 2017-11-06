package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	ex := getMinMaxString("1 2 3 4")
	fmt.Println(ex)
}

func getMinMaxString(s string) (ret []string) {
	slis := strings.Split(s, " ")
	sort.Strings(slis)
	ret = append(ret, slis[0], slis[len(slis)-1])
	return
}
