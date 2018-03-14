#include <iostream>

using namespace std;

int main() {
	int num1, num2, i, num, rem, sum;

	cout << "Enter first number: ";
	cin >> num1;

	cout << "Enter second number: ";
	cin >> num2;

	cout << "Armstrong numbers between " << num1 << " and " << num2 << " are: " << endl;
	for(i = num1; i <= num2; i++) {
		sum = 0;
		num = i;

		for(; num > 0; num /= 10) {
			rem = num % 10;
			sum = sum + rem * rem * rem;
		}

		if(sum == i) { cout << i << endl; }
	}

	return 0;
}
