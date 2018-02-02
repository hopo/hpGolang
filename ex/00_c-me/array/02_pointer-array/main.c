#include <stdio.h>

void string_middle(char *text) {
	printf("%c\n", text[0]);
	printf("%c\n", text[1]);
	printf("%c\n", text[2]);
}

int main(void) {
	char *t;
	t = "powe5hell6";

	string_middle(t);

	return 0;
}
