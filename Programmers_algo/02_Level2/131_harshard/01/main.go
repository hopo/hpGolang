package main

import (
	"fmt"
)

func main() {
	ex1 := harshad(18)  // true, 1+8=9, 18%9 == 0
	ex2 := harshad(13)  // false, 1+3=4, 13%4 == 1
	ex3 := harshad(342) // true, , 3+4+2=9 342%9 == 0
	fmt.Println(ex1)
	fmt.Println(ex2)
	fmt.Println(ex3)
}

func harshad(n int) bool {
	sample := split_int(n)
	sum := sum(sample)
	if n%sum == 0 {
		return true
	}
	return false
}

func split_int(num int) []int {
	numlen := 10
	for {
		if num-numlen < 0 {
			numlen /= 10
			break
		}
		numlen *= 10
	}
	var ea int
	var ret []int
	for {
		ea = num / numlen
		ret = append(ret, ea)
		num %= numlen
		numlen /= 10
		if numlen < 1 {
			break
		}
	}
	return ret
}

func sum(isl []int) int {
	var ret int
	for _, v := range isl {
		ret += v
	}
	return ret
}

/*
# https://programmers.co.kr/learn/challenge_codes/131

"""
양의 정수 x가 하샤드 수이려면 x의 자릿수의 합으로 x가 나누어져야 합니다. 예를들어 18의 자릿수 합은 1+8=9이고, 18은 9로 나누어 떨어지므로 18은 하샤드 수입니다.

Harshad함수는 양의 정수 n을 매개변수로 입력받습니다. 이 n이 하샤드수인지 아닌지 판단하는 함수를 완성하세요.
예를들어 n이 10, 12, 18이면 True를 리턴 11, 13이면 False를 리턴하면 됩니다.
"""

def Harshad(n):
    # n은 하샤드 수 인가요?

    return True

# 아래는 테스트로 출력해 보기 위한 코드입니다.
print(Harshad(18))
*/
