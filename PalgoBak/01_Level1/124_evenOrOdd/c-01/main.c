#include <stdio.h>

char evenOrOdd(int);

int main() {
	char ex1 = evenOrOdd(3); // O
	printf("%c", ex1);

	char ex2 = evenOrOdd(2); // E
	printf("\n%c", ex2);

	return 0;
}

char evenOrOdd(int num) {
	char ret;
	if(num%2 == 0) { ret = 'E'; }
	else { ret = 'O'; }

	return ret;
}
