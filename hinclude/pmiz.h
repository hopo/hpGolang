// **************************************
// *** programiz.com cpp examples api ***
// **************************************


#ifndef PMIZ_H

// Standard Library
#include <vector>

using namespace std;


// method list this header
int reverse_number(int);
int pmiz_gcd(int, int);
int pmiz_lcm(int, int);
bool prime(int); 
vector<int> factor(int n);


// reverse_number()
// p : int
// r : int
// ex) 1234 -> 4321
int reverse_number(int n) {
	int reversedNumber = 0, remainder;

	while(n != 0) {
		remainder = n % 10;
		reversedNumber = reversedNumber * 10 + remainder;
		n /= 10;
	}

	return reversedNumber;
}


// pmiz_gcd() 
// p : int, int
// r : int
// gcd(hcf) n1 and n2
int pmiz_gcd(int n1, int n2) {
	// bigger - smaller ?
	while(n1 != n2) {
		if(n1 > n2) { n1 -= n2; }
		else { n2 -= n1; }
	} 

	return n1;
}


// pmiz_lcm()
// p : int, int
// r : int
// lcm n1 and n2
int pmiz_lcm(int n1, int n2) {
	int max;
	// maximum value between n1 and n2 is stored in max
	max = (n1 > n2) ? n1 : n2;

	do {
		if(max % n1 == 0 && max % n2 == 0) { break; }
		else { ++max; }
	}
	while(true);

	return max;
}


// prime()
// p : int
// r : bool
// prime number
bool prime(int n) {
	int i;
	bool isPrime = true;

	for(i = 2; i <= n / 2; ++i) {
		if(n % i == 0) {
			isPrime = false;
			break;
		}
	}

	return isPrime;
}


// factor()
// p : int
// r : vector<int>
// denominator? factor?
vector<int> factor(int n) {
	int i;
	vector<int> ret;

	for(i = 1; i <= n; ++i) {
		if(n % i == 0) { ret.push_back(i); }
	}

	return ret;
}


#endif
