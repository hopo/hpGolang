// https://programmers.co.kr/learn/challenge_codes/147

#include <iostream>
#include <vector>

using namespace std;

vector<int> fibonacci(int n) {
	long long a, b, c;
	a = 0;
	b = 1;
	c = a + b;
	vector<int> ctnr;
	for (int i = 0; i < n-1 ; i++) {
		a = b;
		b = c;
		ctnr.push_back(b);
		c = a + b;
	}

	return ctnr;
}

int main() {
	int testCase = 10; // 55
	vector<int> r = fibonacci(testCase);
	cout << r[testCase-2] << endl;

	return 0;
}
