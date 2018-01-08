#include <stdio.h>

char* reverseStr(char*);

int main() {
	char* ex1 = reverseStr("Zbcdefg");     // "gfedcbZ"
	int i = 0;
	char prt[100];
	while(ex1[i] != '\0') {
		prt[i] = ex1[i];		
		i++;
	}
	prt[i] = '\0';
	printf("%s", prt);

	return 0;
}

char* reverseStr(char* str) {
	int lnth = 0;

	while(*(str+lnth) != '\0') { lnth++; }
	char rev[lnth];
	
	int i, j;
	for(i = 0; i < lnth; i++) {
		j = lnth-1-i;
		rev[i] = *(str+j);
	}

	char* dt;
	dt = rev;

	return dt;
}
