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
		if (b > n) { break; }
		ctnr.push_back(b);
		c = a + b;
	}

	return ctnr;
}

int main() {
	int testCase = 47; // 8
	cout << fibonacci(testCase).size() << endl;

	return 0;
}
