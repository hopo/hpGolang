// without using temporary variables

#include <iostream>

using namespace std;

int main() {
	int a = 5, b = 10;

	cout << "Before swapping." << endl;
	cout << "a = " << a << ", b = " << b << endl;
	
	/*
	a = a + b; // 15 = 5 + 10, a += b
	b = a - b; // 5 = 15 - 10, b = a-b
	a = a - b; // 10 = 15 - 5, a -= b
	*/

	a += b; b = a-b; a -= b;

	cout << "\nAfter swapping." << endl;
	cout << "a = " << a << ", b = " << b << endl;

	return 0;
}
