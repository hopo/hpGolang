#include <stdio.h>

int main() {
	int c;
	printf("geeks for %ngeeks ", &c);
	printf("%d", c);
	getchar();

	return 0;
}

/*
geeks for geeks 10
*/
