#include <stdio.h>

char *string_middle(char *t);

int main() {
	char text1[] = "power";	// w
	char *ex1 = string_middle(text1);
	printf("%c", *ex1);

	char text2[] = "test";	// es
	char *ex2 = string_middle(text2);
	printf("\n%c%c", *ex2, *(ex2+1));

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

	char *dt;
	dt = ret;

	return dt;
}
