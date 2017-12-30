// https://programmers.co.kr/learn/challenge_codes/146

#include <stdio.h>

int sumDivisor(int);

int main() {
	int ex1 = sumDivisor(10); // 18, [1, 2, 5, 10] 
	printf("%d\n", ex1);

	int ex2 = sumDivisor(12); // 28, [1, 2, 3, 4, 6, 12] 
	printf("%d\n", ex2);

	return 0;
}

int sumDivisor(int n) {
	int ret = 0;	
	for(int i = 1; i < n+1; i++){
		if(n % i == 0) { ret += i; }
	}

	return ret;
}
