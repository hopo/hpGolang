#include <stdio.h>
#include <stdlib.h>

int strToInt(char s[]);

int main() {
	char str1[] = "-1234";
	int ex1 = strToInt(str1);	// -1234
	printf("%d", ex1);

	char str2[] = "392";
	int ex2 = strToInt(str2);	// 392
	printf("\n%d", ex2);

	return 0;
}

int strToInt(char s[]) {
	return atoi(s);
}
