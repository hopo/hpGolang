#include <stdio.h>

int main() {
	char* str[] = {"Hello", "World"};

	printf("str[0]: %s", str[0]);	//Hello
	printf("\nstr[1]: %s", str[1]);	//World
	
	printf("\nsizeof(str): %lu", sizeof(str));	//16
	printf("\nsizeof(str[0]): %lu", sizeof(str[0]));	//8

	return 0;
}
