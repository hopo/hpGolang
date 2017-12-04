// https://programmers.co.kr/learn/challenge_codes/147

#include<iostream>
using namespace std;

long long fibonacci(int n)
{
	long long a, b, c;
	a = 0;
	b = 1;
	c = a + b;
	for (int i = 0; i < n-1 ; i++) {
		a = b;
		b = c;
		c = a + b;
	}
	return b;
}

int main()
{
	int testCase = 47; // 8
	long long testAnswer = fibonacci(testCase);

	cout<<testAnswer;
}
