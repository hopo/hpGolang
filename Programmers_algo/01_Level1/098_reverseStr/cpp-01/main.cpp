#include <iostream>

using namespace std;

string reverseStr(string);

int main() {
	string ex1 = "Zbcdefg";     // "gfedcbZ"
	cout << reverseStr(ex1) << endl;
	
	string ex2 = "deSadFklJfg"; // "gfJlkFdaSed"
	cout << reverseStr(ex2) << endl;

	return 0;
}

string reverseStr(string s) {
	int size;
	while(1){
		if(s[size] == 0) { break;}
		size++;
	}

	int j;
	string ret;
	for(int i = 0; i < size; ++i) {
		j = size-1-i;
		ret += s[j];
	}

	return ret;
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
