#include <stdio.h>

char string_middle(char* t[]) {
	printf("%c\n", t[0]);
	printf("%c\n", t[1]);
	printf("%c\n", t[2]);
	
	return '/';
}

int main(int argc, char* argv[]) {
	char text1[] = "powe5hell6";	// w
	char ex1 = string_middle(text1);
	printf("%c", ex1);

	return 0;
}
