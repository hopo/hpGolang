#include<iostream>
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
	int testCase = 10;
	int testAnswer = sumDivisor(testCase);

	cout << testAnswer;
}
