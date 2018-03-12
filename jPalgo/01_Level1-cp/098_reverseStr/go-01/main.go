package main

import (
	"fmt"
)

func main() {
	ex1 := reverseStr("Zbcdefg")     // "gfedcbZ"
	ex2 := reverseStr("deSadFklJfg") // "lkgfeddaSJF"
	fmt.Println(ex1)
	fmt.Println(ex2)
}

func reverseStr(s string) (ret string) {
	for i := 0; i < 200; i++ {
		for _, v := range s {
			if i == int(v) {
				ret = string(v) + ret
			}
		}
	}
	return
}

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
