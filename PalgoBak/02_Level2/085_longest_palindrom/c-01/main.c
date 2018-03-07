#include <stdio.h>

int longest_palindrom(char *);

int main() {
	char text1[] = "tomotmomtomot";
	int ex1;
	ex1 = longest_palindrom(text1);	// 13
	printf("%d\n", ex1);

	char text2[] = "tomotmkdir";
	int ex2;
	ex2 = longest_palindrom(text2);	// 5
	printf("%d\n", ex2);

	char text3[] = "xmmxaj";
	int ex3;
	ex3 = longest_palindrom(text3);	// 4
	printf("%d", ex3);

	return 0;
}


int longest_palindrom(char *text) {
	int lnth, e, i;
	lnth = e = i = 0;
	while(text[lnth] != '\0') { lnth++; }

	for(e = lnth-1; e > -1; e--) {
		if(text[i] == text[e]) { i++; }
		else { i = 0; }
	}

	return i;
}
