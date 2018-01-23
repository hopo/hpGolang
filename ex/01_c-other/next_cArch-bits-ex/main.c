#include <stdio.h>
#include <stdlib.h>

/*
*minus
	x-n = x+(~n+1)

*mask
	x&0 = 0
	x&~0 = x
	x|0 = x
	x|~0 = ~0
*/

void bprt(int d, int f) {
	if(f != 0) { printf("%d ->\t", d); }
	if(d > 255) { printf("OVERFLOW! "); }
	char str[] = "00000000"; 

	int i = 7;
	while(i >= 0) {
		if(d%2) { str[i] = '1';	}
		--i;
		d = d/2;
	}

	printf("%s", str);
}

void addOK();	// is it overflow??
void fitsBits();	// 변수x가n비트로표현가능한가?
void minusOne();	// '-1' return.
void getByte();	//(3)	
void allEvenBits();	//(2)
void bitxor();	//(1)

int main() {
	//getByte();
	//allEvenBits();
	//bitxor();

	return 0;
}


/* (3)Var x, n' bit is 1 or 0 */
void getByte() {
	int x, n, i = 0;
	x = 0x39;
	n = 5;

	bprt(x, 0); printf(" n: %d", n); // for check.

	while(1) {	
		if(i == n) {
			printf("\n%d", x%2);
			break;
		}
		x /= 2;
		i++;
	}
}


/* (2)All even bits is '1' == TRUE */
void allEvenBits() {
	printf("allEvenBits() ;\n");
	int a, i, flag;
	a = 0xf8;	// 2 8 32 128 .. (mix) is True
	bprt(a, 1);	// for check. binary print
	
	i = 0;
	while(2<<i < a+1) {
		flag = (2<<i)&a;
		i += 2;
		if(flag == 0) {
			printf("\nFALSE");
			exit(0);
		}
	}

	printf("\nTRUE");
}


/* (1)Xor operator of two int */
void bitxor() {
	printf("bitxor() ;\n");

	int a, b;
	a = 0x02;
	b = 0x03;

	printf("%d", (a|b)+(~(a&b)+1));	// using complement
	printf(" %d ", a^b);	// ^ operator
}
