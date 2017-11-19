package main

import (
	"fmt"
)

/*
1 2
1 2 3 6
1 2 4 8
1 2 7 14

smpl: 1344
*/

func main() {
	ex1 := nlcm([]int{2, 6, 8, 14})
	//ex2 := nlcm([]int{2, 6, 8, 14}) // 168
	fmt.Println(ex1)
	//fmt.Println(ex2)
}

func nlcm(nums []int) int {
	smpl := 1
	for _, v := range nums {
		smpl *= v
	}
	fmt.Println("smpl:", smpl)
	return -1
}

/*
# https://programmers.co.kr/learn/challenge_codes/29

"""
두 수의 최소공배수(Least Common Multiple)란 입력된 두 수의 배수 중 공통이 되는 가장 작은 숫자를 의미합니다. 예를 들어 2와 7의 최소공배수는 14가 됩니다. 정의를 확장해서, n개의 수의 최소공배수는 n 개의 수들의 배수 중 공통이 되는 가장 작은 숫자가 됩니다. nlcm 함수를 통해 n개의 숫자가 입력되었을 때, 최소공배수를 반환해 주세요. 예를들어 [2,6,8,14] 가 입력된다면 168을 반환해 주면 됩니다.
"""

def nlcm(num):
    answer = 0

    return answer

# 아래는 테스트로 출력해 보기 위한 코드입니다.
print(nlcm([2,6,8,14]));
*/
