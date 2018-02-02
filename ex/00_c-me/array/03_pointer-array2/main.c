#include <stdio.h>
#include <string.h>

char string_middle(char* t) {
	printf("%zu\n", strlen(t));
	
	return '/';
}

int main(int argc, char* argv[]) {
	char text1[] = "hello";	// w
	char ex1 = string_middle(text1);
	printf("%c", ex1);

	return 0;
}

