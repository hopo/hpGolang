#include <stdio.h>

typedef struct IntArray
{
	int length;
	int *array;
} iarr;

int main() {
	iarr ia;
	ia.array[0] = 1;
	ia.array[1] = 3;
	ia.array[2] = 5;
	ia.array[3] = 7;
	ia.array[4] = 9;
	ia.length = 4;

	printf("%d", ia.array[0]);
	printf(" %d", ia.array[1]);
	printf(" %d", ia.array[2]);
	printf("\nsizeof: %lu", sizeof(ia.array));

	return 0;	
}
