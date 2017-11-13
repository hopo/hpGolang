// another solution?
package main

import (
	"fmt"
)

func main() {
	ex1 := collats(6)  // 8
	ex2 := collats(11) // 14
	fmt.Println(ex1)
	fmt.Println(ex2)
}

func collats(num int) int {
	chk := []int{num}
	var ret int
	for {
		if ret == 500 {
			ret = -1
			break
		}
		if num == 1 {
			fmt.Println(chk)
			break
		}

		if num%2 == 0 {
			num = num / 2
			chk = append(chk, num)
			ret++
		} else {
			num = num*3 + 1
			chk = append(chk, num)
			ret++
		}
	}
	return ret
}

/*
# https://programmers.co.kr/learn/challenge_codes/14

"""
1937년 Collatz란 사람에 의해 제기된 이 추측은, 입력된 수가 짝수라면 2로 나누고, 홀수라면 3을 곱하고 1을 더한 다음, 결과로 나온 수에 같은 작업을 1이 될 때까지 반복할 경우 모든 수가 1이 된다는 추측입니다. 예를 들어, 입력된 수가 6이라면 6→3→10→5→16→8→4→2→1 이 되어 총 8번 만에 1이 됩니다. collatz 함수를 만들어 입력된 수가 몇 번 만에 1이 되는지 반환해 주세요. 단, 500번을 반복해도 1이 되지 않는다면 –1을 반환해 주세요.
"""

def collatz(num):
    answer = 0

    return answer

# 아래는 테스트로 출력해 보기 위한 코드입니다.
print(collatz(6))
*/
