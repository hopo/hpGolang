#include <stdio.h>
#include <stdlib.h>

const char ST = '*';
const char NL = '\n';
const char ZR = '\0';

void printTriangle(int num, char ret[]);

int main() {
	int num = 4;
	char ret[100];
	printTriangle(num, ret); // *\n**\n***
	printf("%s", ret);

	return 0;
}

void printTriangle(int num, char ret[]) {
	if(num > 12) {
		printf("num is over.");
		exit(0);
	}

	int i, j, r = 0;
	for(i = 1; i < num+1; i++) {
		for(j = 0; j < i; j++) {
			ret[r] = ST;
			r++;
		}
		ret[r] = NL;
		r++;
	}
	ret[r] = ZR;
}
