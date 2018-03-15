#include <stdio.h>

int sumDivisor(int);

int main() {
	int ex1;
	ex1 = sumDivisor(10); // 18, [1, 2, 5, 10] 
	printf("%d", ex1);

	int ex2;
	ex2 = sumDivisor(12); // 28, [1, 2, 3, 4, 6, 12] 
	printf("\n%d", ex2);

	return 0;
}

int sumDivisor(int n) {
	int ret = 0;	

	/* Yaksoo of n. total sum */ 
	for(int i = 1; i < n+1; i++){
		if(n % i == 0) { ret += i; }
	}

	return ret;
}
