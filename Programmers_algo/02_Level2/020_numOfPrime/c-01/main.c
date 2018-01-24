#include <stdio.h>

int numOfPrime(int);

int main() {
	int num1, ex1;
	num1 = 10;
	ex1 = numOfPrime(num1);	// 4, [2 3 5 7]
	printf("%d\n", ex1);

	int num2, ex2;
	num2 = 5;
	ex2 = numOfPrime(num2);	// 3, [2 3 5]
	printf("%d", ex2);
	
	return 0;
}


int numOfPrime(int num) {
	int i, j, count;
	count = 1; // counting about "2", 2 is prime number.

	for(i = 3; i < num+1; i++) {
		for(j = 2; j < i; j++) {
			if(i%j == 0) { break; }	// i is not prime.
			if(i-j == 1) { count++; }	// (j) loop finish.
		}
	}

	return count;
}
