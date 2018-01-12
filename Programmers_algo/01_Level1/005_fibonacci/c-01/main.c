#include <stdio.h>

long long fibonacci(int);

int main() {
	int ex1;
	ex1 = fibonacci(10); // 55
	printf("%d", ex1);

	return 0;
}

long long fibonacci(int n) {
	long long a, b, c, i;
	a = 0;
	b = 1;
	c = a+b;
	for(i = 0; i < n-1; i++) {
		a = b;
		b = c;
		c = a+b;
	}

	return b;
}
