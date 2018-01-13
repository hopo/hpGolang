#include <stdio.h>

int main() {
	int x= 0x0001;
	int y= 0x000e;

	int z = x+y;
	printf("0x%.4x", z);
}
