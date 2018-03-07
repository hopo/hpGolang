#include <stdio.h>

typedef struct Data
{
	char c1;
	char c2;
} data;

data string_middle(char text[]);

int main() {
	char text1[] = "power";	// w
	data ex1 = string_middle(text1);
	printf("%c", ex1.c1);

	char text2[] = "test";	// es
	data ex2 = string_middle(text2);
	printf("\n%c%c", ex2.c1, ex2.c2);

	return 0;
}

data string_middle(char txt[]) {
	int i, lnth;
	for(i = 0; txt[i] != '\0'; i++);
	lnth = i;
	
	data ret;	
	switch(lnth%2) {
		case(1):
			ret.c1 = txt[lnth/2];
			break;
		case(0):
			ret.c1 = txt[lnth/2-1];
			ret.c2 = txt[lnth/2];
			break;
	}

	return ret;
}
