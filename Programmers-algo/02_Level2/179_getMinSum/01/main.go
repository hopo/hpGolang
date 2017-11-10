package main

import (
	"fmt"
)

func main() {
	ex1 := getMinMax([]int{1, 2}, []int{3, 4}) // 10
	fmt.Println(ex1)

}

func getMinMax(a, b []int) int

/*
// https://programmers.co.kr/learn/challenge_codes/179

#include<iostream>
#include<vector>
using namespace std;

int getMinSum(vector<int> A, vector<int> B)
{
	int answer = 0;

	return answer;
}
int main()
{
	vector<int> tA{1,2}, tB{3,4};

	//아래는 테스트 출력을 위한 코드입니다.
	cout<<getMinSum(tA,tB);
}
*/
