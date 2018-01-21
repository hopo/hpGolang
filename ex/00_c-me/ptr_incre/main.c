#include <stdio.h>

int main() {
	int d;
	d = 42;

	printf("d: %d\n", d);	

	int *ptr_d;
	ptr_d = &d;

	printf("ptr_d: %p\n", ptr_d);	
	printf("*ptr_d: %d\n", *ptr_d);
	
	int *ptr_d1;
	ptr_d1 = ptr_d+1;

	printf("ptr_d1: %p\n", ptr_d1);
	printf("*ptr_d1: %d\n", *ptr_d1);
	
	int *ptr_d2;
	ptr_d2 = ptr_d+2;

	printf("ptr_d2: %p\n", ptr_d2);
	printf("*ptr_d2: %d\n", *ptr_d2);

	return 0;
}
