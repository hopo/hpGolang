#include <stdio.h>

int main() {
	int x = 40000;
	int y = 50000;
	float z = 1e16;

	printf("%p", &x);
	printf("\n%d", x*x);
	printf("\n%d", y*y);
	printf("\n%f", z);
	
	return 0;
}
