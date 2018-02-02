#include <stdio.h>

void prnt(char* carr);

int main() {
	char carr[2];

	carr[0] = 'h';
	carr[1] = 'i';

	prnt(carr);

	return 0;
}

void prnt(char* carr) {
	printf("%c%c", carr[0], carr[1]);
}
