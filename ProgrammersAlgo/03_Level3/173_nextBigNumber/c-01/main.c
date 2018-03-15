#include <stdio.h>

int nextBigNumber(int);
int oneCount(int);

int main() {
	int n1 = 78, ex1;
	ex1 = nextBigNumber(n1); // 83
	printf("%d\n", ex1);

	int n2 = 47, ex2;
	ex2 = nextBigNumber(n2); // 55
	printf("%d", ex2);
}

int nextBigNumber(int n) {
	int nb, ib, i = n;
	nb = oneCount(n);

	while(nb != ib) {
		i++;
		ib = oneCount(i);
	}

	return i;
}

int oneCount(int num) {
	int b = 0;
	while(num != 0) {
		if(num%2){ b++; }
		num /= 2;
	}

	return b;
}
