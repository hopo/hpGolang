#include <stdio.h>

char* string_middle(char* t);

int main(int argc, char* argv[]) {
	char text1[] = "power";	// w
	char* ex1 = string_middle(&text1);
	printf("%c", *ex1);

	printf("\n");

	char text2[] = "test";	// es
	char* ex2 = string_middle(&text2);
	printf("%c%c", *ex2, *(ex2+1));

	return 0;
}

char* string_middle(char *t) {
	int i, len = 0;
	for(i = 0; t[i] != 0; i++) { len++; }
	
	char* ret;
	switch(len%2) {
		case(1):
			*ret = t[len/2];
			break;
		case(0):
			*ret = t[len/2-1];
			*(ret+1) = t[len/2];
			break;
	}

	return ret;
}
