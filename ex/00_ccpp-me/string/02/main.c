#include <stdio.h>

int main() {
	char str[] = "Hello World";

	printf("str[]: %s", str);	//Hello World
	printf("\nstr[0]: %c", str[0]);	//H
	
	printf("\nsizeof(str): %lu", sizeof(str));	//12
	printf("\nsizeof(str[0]): %lu", sizeof(str[0]));	//1

	printf("\nstr[11]: %hhd", str[11]);	//0

	return 0;
}
