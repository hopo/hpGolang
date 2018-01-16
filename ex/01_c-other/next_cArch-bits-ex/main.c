#include <stdio.h>

/*
!minus
	x-n = x+(~n+1)

!mask
	x&0 = 0
	x&~0 = x
	x|0 = x
	x|~0 = ~0
*/

void bitxor(int, int);	// xor operator of two int
void allEvenBits(int, int);	// 짝수번째비트가'1'인가?
void getByte(int, int);	//	변수x, n번째의 byte?
void minusOne(int, int);	// '-1' return.
void fitsBits(int, int);	// 변수x가n비트로표현가능한가?
void addOK(int, int);	// is it overflow??

int main() {
	int a = 0x04;
	int b = 0x03;
	bitxor(a, b);

	return 0;
}

void bitxor(int a, int b) {
	printf("%d", (a|b)+(~(a&b)+1));	// using complement
	printf(" %d ", a^b);	// ^ operator
}
