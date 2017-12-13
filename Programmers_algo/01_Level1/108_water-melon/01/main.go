package main

import (
	"fmt"
)

func main() {
	ex1 := water_melon(3)
	ex2 := water_melon(4)
	fmt.Println(ex1) // "WMW"
	fmt.Println(ex2) // "WMWM"
}

func water_melon(n int) (ret string) {
	for i := 1; i < n+1; i++ {
		if i%2 == 1 {
			ret += "W"
		} else {
			ret += "M"
		}
	}
	return
}

/*
# https://programmers.co.kr/learn/challenge_codes/108

def water_melon(n):
    # 함수를 완성하세요.
    return ""


# 실행을 위한 테스트코드입니다.
print("n이 3인 경우: " + water_melon(3));
print("n이 4인 경우: " + water_melon(4));
*/
