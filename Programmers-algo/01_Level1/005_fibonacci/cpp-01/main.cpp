// https://programmers.co.kr/learn/challenge_codes/147

#include <iostream>

using namespace std;

long long fibonacci(int n) {
	long long a, b, c, i;
	a = 0;
	b = 1;
	c = a + b;
	for (i = 0;i < n-1; i++) {
		a = b;
		b = c;
		c = a + b;
	}
	return b;
}

int main() {
	int testCase = 10; // 55
	cout <<  fibonacci(testCase) <<  endl;

	return 0;
}
