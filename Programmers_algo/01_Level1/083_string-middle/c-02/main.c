#include <stdio.h>

char *string_middle(char text[]);

int main() {
	char text1[] = "power";	// w
	char *ex1 = string_middle(text1);
	printf("%c", ex1[0]);

	char text2[] = "test";	// es
	char *ex2 = string_middle(text2);
	printf("\n%c%c", ex2[0], ex2[1]);

	return 0;
}

char *string_middle(char txt[]) {
	int i, lnth;
	for(i = 0; txt[i] != '\0'; i++);
	lnth = i;
	
	char ret[3];
	switch(lnth%2) {
		case(1):
			ret[0] = txt[lnth/2];
			ret[1] = '\0';
			break;
		case(0):
			ret[0] = txt[lnth/2-1];
			ret[1] = txt[lnth/2];
			ret[2] = '\0';
			break;
	}

	char *dt;
	dt = ret;

	return dt;
}
