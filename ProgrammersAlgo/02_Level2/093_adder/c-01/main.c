#include <stdio.h>

int adder(int a, int b);

int main() {
	int a = 3, b = 5, ex;
	ex = adder(a, b);	// 12, 3+4+5
	printf("%d", ex);

	return 0;
}

int adder(int a, int b) {
	if(a == b) { return a; }

	int tmp;
	if(a > b) {
		tmp = a;
		a = b;
		b = tmp;
	}

	int i, ret;
	for(i = a; i < b+1; i++) { ret += i; }

	return ret;
}
