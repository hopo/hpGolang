#include <stdio.h>

char* printTriangle(int);

int main() {
	char* ex1 = printTriangle(3); // *\n**\n***
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

char* printTriangle(int num) {
	int i, j, r = 0;
	char ret[100];
	for(i = 1; i < num+1; i++) {
		for(j = 0; j < i; j++) {
			ret[r] = '*';
			r++;
		}
		ret[r] = '\n';
		r++;
	}
	ret[r] = '\0';

	char* dt = ret;

	return dt;
}
