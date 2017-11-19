package main

import (
	"fmt"
)

func main() {
	ex1 := jumpCase(5) // 5
	fmt.Println(ex1)
}

func jumpCase(n int) int {
	fmt.Println("n:", n)

	isl := big_dduim(n)
	fmt.Println("isl:", isl)

	ssl := make_vari_ssl(isl)
	fmt.Println("ssl:", ssl)

	return -1
}

func big_dduim(n int) []int {
	var isl []int
	for {
		if n == 2 {
			isl = append(isl, n)
			break
		}
		if n == 1 {
			isl = append(isl, n)
			break
		}
		if n-2 > 0 {
			isl = append(isl, 2)
			n -= 2
		}
	}
	return isl
}

func make_vari_ssl(isl []int) []string {
	var s string
	var ssl []string
	for i, _ := range isl {
		for j, _ := range isl {
			if i != j && isl[i] != isl[j] {
				isl[i], isl[j] = isl[j], isl[i]
				for _, v := range isl {
					s += string(v + 48)
				}
				ssl = append(ssl, s)
				s = ""
				isl[j], isl[i] = isl[i], isl[j]
			}
		}
	}
	for i, _ := range ssl {
		for j, _ := range ssl {
			if i != j && ssl[i] == ssl[j] {
				fmt.Println("D!")
			}
		}
	}
	return ssl
}
