#include <stdio.h>

typedef int Iarr[];

int main() {
	Iarr x = {};
	x[0]= 2;
	x[1]= 4;
	x[2]= 6;

	printf("%d", x[0]);
	printf(" %d", x[1]);
	printf(" %d", x[2]);


	printf("\nsize: %lu", sizeof(x));
	printf("\nsize2: %lu", sizeof(x[0]));

	return 0;
}
