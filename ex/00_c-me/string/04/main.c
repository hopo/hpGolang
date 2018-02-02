#include <stdio.h>

void strchk(char** str) {
	
	printf("*str: %s\n", *str);	//Hello
	printf("*(str+1): %s\n", *(str+1));	//World

	printf("sizeof(str): %lu\n", sizeof(str));	//8
	printf("sizeof(str[0]): %lu\n", sizeof(str[0]));	//8

	printf("ptr str[0]: %p\n", *str);
	printf("ptr str[1]: %p", *(str+1));

}

int main() {
	char* str[] = {"Hello", "World"};

	strchk(str);

	return 0;
}
