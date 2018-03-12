package main

import (
	"fmt"
)

func main() {
	ex1 := string_middle("power") // [w]
	ex2 := string_middle("test")  // [e s]
	fmt.Println(ex1)
	fmt.Println(ex2)
}

func string_middle(s string) (ret []string) {
	l := len(s)
	c := l / 2
	if l%2 == 0 {
		r2 := string(s[c-1])
		ret = append(ret, r2)
	}
	r1 := string(s[c])
	ret = append(ret, r1)
	return
}

/*
# https://programmers.co.kr/learn/challenge_codes/83

def string_middle(str):
    # 함수를 완성하세요

    return ""

# 아래는 테스트로 출력해 보기 위한 코드입니다.
print(string_middle("power"))
*/
