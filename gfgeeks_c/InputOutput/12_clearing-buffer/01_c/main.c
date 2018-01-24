#include <stdio.h>

int main() {
	char str[80], ch;
	
	// Scan input from user - GeeksforGeeks for example
	scanf("%s", str);

	// flushes the standard input (clear the input buffer)
	/*
	while((getchar()) != '\n');
	*/

	// Scan character from user - 'a' for example
	ch = getchar();

	// Printing character array, prints "GeeksforGeeks"
	printf("%s\n", str);

	// This does not print character 'a'
	printf("%c", ch);

	return 0;
}
