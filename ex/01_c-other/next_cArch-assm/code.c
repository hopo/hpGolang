// gcc -c
// gcc -s
// objdump -d code.o
// gdb /* gcc debugger */

#include <stdio.h>

int accum = 0;

int sum(int x, int y) {
	int t = x + y;
	accum += t;

	return t;
}

int main() {
	return sum(2, 3);
}
