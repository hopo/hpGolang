package main

import (
	"fmt"
)

func main() {
	ex1 := numPY("pPoooyY")
	ex2 := numPY("Pyy")
	fmt.Println(ex1)
	fmt.Println(ex2)
}

func numPY(s string) bool {
	return false
}

/*
# https://programmers.co.kr/learn/challenge_codes/97

def numPY(s):
    # 함수를 완성하세요

    return True



# 아래는 테스트로 출력해 보기 위한 코드입니다.
print( numPY("pPoooyY") )
print( numPY("Pyy") )
*/
