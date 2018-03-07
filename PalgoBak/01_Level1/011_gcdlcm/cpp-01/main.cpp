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

// Great Common Denominator a and b
int gcd(int a, int b) {
	auto ad = denom(a);	
	auto bd = denom(b);	

	int max;
	for (int i = 0; i < ad.size(); i++) {
		for (int j = 0; j < bd.size(); j++) {
			if (ad[i] == bd[j]) { max = ad[i]; }
		}
	}
	return max;
}


// Lower Common Multiple a and b
int lcm(int a, int b) {
	return a * b / gcd(a, b); // maybe nee edit
}

// gcdlcm a and b
vector<int> gcdlcm(int a, int b) {
	int g = gcd(a, b);
	int l = lcm(a, b);
	vector<int> answer = {g, l};

	return answer;
}

int main() {
	int a = 3, b = 12; // {3, 12}
	vector<int> testAnswer = gcdlcm(a, b);
	cout << testAnswer[0] << " " << testAnswer[1];
	
	return 0;
}
