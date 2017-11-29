#include <stdio.h>
#include <limits.h>

int main() {
	printf("Storage size for int : %d \n", sizeof(int));
	printf("Storage size for float : %d \n", sizeof(float));
   	printf("Minimum float positive value: %E\n", FLT_MIN );
  	printf("Maximum float positive value: %E\n", FLT_MAX );
   	
	return 0;
}
