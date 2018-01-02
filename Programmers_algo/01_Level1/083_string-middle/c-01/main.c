#include <stdio.h>

char *string_middle(char *);

int main() {
	char text1[] = "power";	// w
	char *ex1 = string_middle(&text1);
	printf("%c", *ex1);

	printf("\n");

	char text2[] = "test";	//es
	char *ex2 = string_middle(&text2);
	printf("%c%c", *ex2, *(ex2+1));

	return 0;
}


char *string_middle(char *t) {
	int i, len = 0;
	for(i = 0; t[i] != 0; i++) len++;
	
	char ret[2];
	if(len&1) { ret[0] = t[len/2]; }
	else {
		ret[0] = t[len/2-1];
		ret[1] = t[len/2];
	}

	return &ret;
}
