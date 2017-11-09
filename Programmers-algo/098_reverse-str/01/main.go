package main

import (
	"fmt"
)

func main() {
	ex := reverseStr("Zbcdefg") // "gfedcbZ"
	fmt.Println(ex)
}

func reverseStr(s string) (ret string)

/*
// https://programmers.co.kr/learn/challenge_codes/98

public class ReverseStr {
	public String reverseStr(String str){

		return "";
	}

	// 아래는 테스트로 출력해 보기 위한 코드입니다.
	public static void main(String[] args) {
		ReverseStr rs = new ReverseStr();
		System.out.println( rs.reverseStr("Zbcdefg") );
	}
}
*/
