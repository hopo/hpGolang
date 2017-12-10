//...ing
//https://programmers.co.kr/learn/challenge_codes/149

#include<vector>
#include<iostream>

using namespace std;


// vector int print like array type
void vctrprt(vector<int> vctr) {
	cout << "{";
	for (int i = 0; i < vctr.size(); i++) {
		if (i == vctr.size()-1) {
			cout << vctr[i];
			break;
		}
		cout << vctr[i] << ", ";
	}
	cout << "}";
}

// denominator 1 ~ n
vector<int> denom(int d) {
	vector<int> box;
	for (int i = 1; i < d+1; i++) {
		if (d % i == 0) {
			box.push_back(i);
		}
	}
	return box;
}


/*
vector<int> gcdlcm(int a, int b) {
	vector<int> answer;

	return answer;
}
*/

int main() {
	int a=3, b=12; // [3, 12]
	// vector<int> testAnswer = gcdlcm(a, b);
	// cout << testAnswer[0] << " " << testAnswer[1];
	
	// test
	auto dv = denom(12);
	vctrprt(dv); // checker
}
