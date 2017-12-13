#ifndef PMIZ_H

// programiz.com

#include <iostream>

using namespace std;


// method list
int reverse_number(int);


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


#endif
