// https://programmers.co.kr/learn/challenge_codes/146

#include <iostream>

using namespace std;

int sumDivisor(int n) {
	int ret = 0;	
	for(int i = 1; i < n+1; i++){
		if(n % i == 0) {
			ret += i;
		}
	}
	return ret;
}

int main() {
	int num1 = 10; // 18, [1, 2, 5, 10] 
	cout << sumDivisor(num1) << endl; 

	int num2 = 12; // 28, [1, 2, 3, 4, 6, 12] 
	cout << sumDivisor(num2) << endl;
	return 0;
}
