// https://programmers.co.kr/learn/challenge_codes/146

#include <vector>
#include <iostream>

using namespace std;

vector<int> sumDivisor(int n) {
	vector<int> arr;	
	for(int i = 1; i < n; i++){
		if(n % i == 0) {
			cout << i;
		}
	}
	return arr;
}
int main() {
	int testCase = 10;
	vector<int> testAnswer = sumDivisor(testCase);

	cout << testAnswer;
}
