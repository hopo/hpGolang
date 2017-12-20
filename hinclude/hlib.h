// *************************
// *** For cpp api by hp ***
// *************************


#ifndef HLIB_H

// reference
#include <string>
#include <sstream>
#include <vector>

using namespace std;


// method list this header
template <class T> string boxprt(T);

int gcd(int, int);
int lcm(int, int);
vector<int> factor(int);
vector<int> gcdlcm(int, int);

vector<int> fibonacci(int);


// boxprt()
// p : <class T>
// r : string
// vector int print like array type
template <class T>
string boxprt(T box) {
	stringstream ss; // sstring::stringstream

	ss << "[";
	for (int i = 0; i < box.size(); i++) { 
		if (i == box.size()-1) {
			ss << box[i];
			break;
		}
		ss << box[i] << " ";
	}
	ss << "]";

	return ss.str(); // sstring::stringstream::str()
}


// factor()
// p : int
// r : vector<int> 
// denominator 1 ~ n
vector<int> factor(int d) {
	vector<int> box;

	for (int i = 1; i < d+1; i++) {
		if (d % i == 0) { box.push_back(i); } // vector::push_back()
	}

	return box;
}


// *** gcd()
// p : int, int
// r : int
// Great Common Denominator a and b
int gcd(int a, int b) {
	auto ad = factor(a);	
	auto bd = factor(b);	

	int max;
	for (int i = 0; i < ad.size(); i++) {
		for (int j = 0; j < bd.size(); j++) {
			if (ad[i] == bd[j]) { max = ad[i]; }
		}
	}

	return max;
}


// lcm()
// p : int, int
// r : int
// Lower Common Multiple a and b
int lcm(int a, int b) {
	return a * b / gcd(a, b); // maybe need edit
}


// gcdlcm()
// p : int, int
// r : vector<int>
// gcdlcm a and b
vector<int> gcdlcm(int a, int b) {
	int g = gcd(a, b);
	int l = lcm(a, b);
	vector<int> answer = {g, l};

	return answer;
}


// fibonacci()
// p : int 
// r : vector<int>
// size = n, fibonacci vector int
// vector reference
vector<int> fibonacci(int n) {
	long long a, b, c;
	a = 0;
	b = 1;
	c = a + b;
	vector<int> box;

	for (int i = 0; i < n ; i++) {
		a = b;
		b = c;
		box.push_back(b);
		c = a + b;
	}

	return box;
}


#endif
