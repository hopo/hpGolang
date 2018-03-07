#include <stdio.h>

int jumpCase(int);

int main() {
	int n1 = 4, ex1;
	ex1 = jumpCase(n1); // 5
	printf("%d\n", ex1);

	int n2 = 7, ex2;
	ex2 = jumpCase(n2); // 21
	printf("%d", ex2);
}

int jumpCase(int n) {
	int i, ret, a = 0, b = 1;
	for(i = 0; i < n+1; i++) {
		ret = a + b;
		b = a;
		a = ret;
	}
	
	return ret;
}
