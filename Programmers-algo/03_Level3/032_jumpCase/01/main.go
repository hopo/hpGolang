package main

import (
	"fmt"
)

func main() {
	//ex1 := jumpCase(4) // 5
	ex2 := jumpCase(7)
	//fmt.Println(ex1)
	fmt.Println(ex2)
}

func jumpCase(n int) int {
	fmt.Println("n:", n)
	fmt.Println("fibo:", fibo(n))
	dduim1(n)
	dduim2(n)
	return -1
}

func dduim1(n int) {
	var isl []int
	var x int
	for {
		isl = append(isl, 1)
		x = n - 1
		n = x
		if x == 0 {
			break
		}
	}
	fmt.Println("dduim1:", isl)
}

func dduim2(n int) {
	var isl []int
	var x int
	for {
		isl = append(isl, 2)
		x = n - 2
		n = x
		if x == 0 {
			break
		}
		if x == 1 {
			isl = append(isl, 1)
			break
		}
	}
	fmt.Println("dduim2:", isl)
}

func fibo(n int) int {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return b
}

/*
# https://programmers.co.kr/learn/challenge_codes/32

"""
효진이는 멀리 뛰기를 연습하고 있습니다. 효진이는 한번에 1칸, 또는 2칸을 뛸 수 있습니다. 칸이 총 4개 있을 때, 효진이는
(1칸, 1칸, 1칸, 1칸)
(1칸, 2칸, 1칸)
(1칸, 1칸, 2칸)
(2칸, 1칸, 1칸)
(2칸, 2칸)
의 5가지 방법으로 맨 끝 칸에 도달할 수 있습니다. 멀리뛰기에 사용될 칸의 수 n이 주어질 때, 효진이가 끝에 도달하는 방법이 몇 가지인지 출력하는 jumpCase 함수를 완성하세요. 예를 들어 4가 입력된다면, 5를 반환해 주면 됩니다.
"""

def jumpCase(num):
    answer = 0

    return answer

#아래는 테스트로 출력해 보기 위한 코드입니다.
print(jumpCase(4))
*/
