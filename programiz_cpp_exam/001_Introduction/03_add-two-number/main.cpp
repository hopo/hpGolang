#include <iostream>

using namespace std;

int main() {
	int firstNumber, secondNumber, sumOfNumbers;

	cout << "Enter two integer: ";
	cin >> firstNumber >> secondNumber;

	// sum of two numbers in stored in variable sumOfTwoNumbers
	sumOfNumbers = firstNumber + secondNumber;

	// Prints sum
	cout << firstNumber << " + " << secondNumber << " = " << sumOfNumbers;

	return 0;
}
