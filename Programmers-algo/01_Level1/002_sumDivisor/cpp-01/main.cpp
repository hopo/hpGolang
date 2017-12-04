// https://programmers.co.kr/learn/challenge_codes/146

#include <vector>
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
	int num2 = 12; // 28, [1, 2, 3, 4, 6, 12] 
	int ex1 = sumDivisor(num1);
	int ex2 = sumDivisor(num2);
	cout << ex1 << endl;
	cout << ex2;
	return 0;
}
