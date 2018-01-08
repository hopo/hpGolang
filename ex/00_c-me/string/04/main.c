#include <stdio.h>

void strchk(char** str) {
	
	printf("*str: %s", *str);	//Hello
	printf("\n*(str+1): %s", *(str+1));	//World

	printf("\nsizeof(str): %lu", sizeof(str));	//8
	printf("\nsizeof(str[0]): %lu", sizeof(str[0]));	//8

	printf("\nptr str[0]: %p", *str);
	printf("\nptr str[1]: %p", *(str+1));

}

int main() {
	char* str[] = {"Hello", "World"};

	strchk(str);

	return 0;
}
