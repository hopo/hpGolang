// convert decimal number to binary

#include <iostream>
#include <cmath>

using namespace std;

long long convertDecimalToBinary(int);

int main() {
	int n, binaryNumber;

	cout << "Enter a decimal number: ";
	cin >> n;
	binaryNumber = convertDecimalToBinary(n);
	cout << n << " in decimal = " << binaryNumber << " int binary" << endl;

	return 0;
}

long long convertDecimalToBinary(int n) {
	long long binaryNumber = 0;
	int remainder, i = 1, step = 1;

	while(n !=0) {
		remainder = n%2;
		cout << "Step " << step++ << ":\t";
		cout << n << "/2,\tRemainder = " << remainder;
		cout << ", Quotient = " << n/2 << endl;

		n /= 2;
		binaryNumber += remainder*i;
		i *= 10;
	}

	return binaryNumber;
}
